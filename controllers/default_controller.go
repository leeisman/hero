package controllers

import (
	"github.com/labstack/echo/v4"
	"hero/pkg/logger"
	"net/http"
)

type Response struct {
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseFail(err error, endPoint string, c echo.Context) error {
	logger.Str(endPoint, err.Error()).Msg("")
	response := &Response{
		Msg:    err.Error(),
		Status: false,
		Data:   "",
	}
	return c.JSON(http.StatusInternalServerError, response)
}

func ResponseSuccess(data interface{}, c echo.Context) error {
	response := &Response{
		Msg:    "",
		Status: true,
		Data:   data,
	}
	return c.JSON(http.StatusOK, response)
}
