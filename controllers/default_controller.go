package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"hero/pkg/logger"
	"net/http"
)

type Response struct {
	Msg    string      `json:"msg"`
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func ResponseFail(err error, c echo.Context) error {
	logger.Error(fmt.Sprintf("%s | %s", c.Request().RequestURI, err.Error()))
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
