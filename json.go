package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

const (
	Err401Str = "Invalid login or password"
	Err405Str = "Parameter %s is required"
	Err406Str = "Parameter %s has invalid format"
)

// AccessToDB - интерефейс, дающий возможность работать как через sql.DB, так и через sql.Tx
type AccessToDB interface {
	QueryRow(query string, argv ...interface{}) *sql.Row
	Query(query string, argv ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}

// Response - струтура для стандартного ответа
type Response struct {
	ErrorCode    int         `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	FoundCount   int         `json:"foundCount"`
	Result       interface{} `json:"result"`
}

// DecodeWrapper - обертка над json.Decoder.Decode
func DecodeWrapper(dec *json.Decoder, out interface{}) (int, string) {
	if err := dec.Decode(&out); err != nil {
		switch err.(type) {
		case *json.UnmarshalTypeError:
			return 406, "Параметр " + err.(*json.UnmarshalTypeError).Field + " имеет неверный формат"
		default:
			return 400, "Тело запроса содержит некорректную структуру JSON"
		}
	}

	return 0, ""
}

// DecodeBytes - декодирование []bytes с проверками
func DecodeBytes(in []byte, out interface{}) (int, string) {
	dec := json.NewDecoder(bytes.NewReader(in))
	return DecodeWrapper(dec, out)
}

// DecodeReader - декодирование io.Reader с проверками
func DecodeReader(in io.Reader, out interface{}) (int, string) {
	dec := json.NewDecoder(in)
	return DecodeWrapper(dec, out)
}

// Decode - функция декодирования полученного запроса
//	in м.б. либо []byte, либо io.Reader
func Decode(in interface{}, out interface{}) (int, string) {
	switch in.(type) {
	case []byte:
		return DecodeBytes(in.([]byte), out)
	case io.Reader:
		return DecodeReader(in.(io.Reader), out)
	}
	return -1, ""
}

// Validator - провера заполненности полей
func Validator(in interface{}) (bool, int, string) {
	var (
		err     error
		errCode = 405
	)
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}
		return name
	})

	if err = v.Struct(in); err != nil {
		if err.(validator.ValidationErrors)[0].ActualTag() != "required" {
			errCode = 406
		}

		return false, errCode, strings.Join(strings.Split(err.(validator.ValidationErrors)[0].Namespace(), ".")[1:], ".")
	}

	return true, 0, ""
}

func SuccessResponse(c echo.Context, resp ...interface{}) {
	var response Response
	response.ErrorCode = 0
	response.ErrorMessage = ""
	response.Result = resp[0]
	if len(resp) > 1 {
		response.FoundCount = resp[1].(int)
	}
	c.JSON(http.StatusOK, response)
}

func FatalResponse(c echo.Context, code int, err string) {
	log.Println(err)
	var response Response
	response.ErrorMessage = "Internal Error"
	if code != 500 {
		response.ErrorMessage = err
	}
	response.ErrorCode = code
	c.JSON(http.StatusOK, response)
}

// UnmarshalJSON - анмаршаллит json в структуру
func UnmarshalJSON(in, out interface{}) (int, string) {
	var (
		errCode    int
		errMessage string
		ok         bool
		body       []byte
	)

	switch in.(type) {
	case echo.Context:
		body, _ = ioutil.ReadAll(in.(echo.Context).Request().Body)
	case []byte:
		body = in.([]byte)
	}

	if errCode, errMessage = Decode(body, &out); errCode > 0 {
		return errCode, errMessage
	}
	// валидация
	if ok, errCode, errMessage = Validator(out); !ok {
		tmpl := Err405Str
		if errCode != 405 {
			tmpl = Err406Str
		}
		return errCode, strings.Replace(tmpl, "%s", errMessage, 1)
	}

	return 0, ""
}
