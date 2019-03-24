package foodWay

import "database/sql"

type Config struct {
	Port       string `json:"port"`
	DBUrl      string `json:"dbUrl"`
	UploadPath string `json:"uploadPath"`
	Domain     string `json:"domain"`
}

type accessToDB interface {
	QueryRow(query string, argv ...interface{}) *sql.Row
	Query(query string, argv ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}
