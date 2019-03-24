package main

import (
	"hack/foodWay"

	"github.com/labstack/echo"
)

func GetAllMallsHandler(c echo.Context) error {
	var (
		err   error
		malls []producter.Point
	)
	if malls, err = foodWay.GetMalls(db); err != nil {
		FatalResponse(c, FatalErrorCode, err.Error())
		return nil
	}
	SuccessResponse(c, malls)
	return nil
}
