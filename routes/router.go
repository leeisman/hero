package routes

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hero/configs"
	"hero/pkg/db/mysql"
	"hero/pkg/logger"
	"hero/utils"
	"io/ioutil"
	"os"
)

func Run() error {
	port := configs.Get("server.port")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	if true {
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			logger.Print(c.Request().URL, "BodyDump reqBody", requestID, string(reqBody))
			logger.Print(c.Request().URL, "BodyDump resBody", requestID, string(resBody))
		}))
		//安全性驗證
		e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {

			if c.Request().Method == "GET" {
				key = "Xn2r5u8x"
				return true, nil
			}

			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			body := &struct {
				FbUserID string `json:"fb_user_id"`
			}{}
			if c.Request().Body != nil { // Read
				reqBody, _ := ioutil.ReadAll(c.Request().Body)
				err := json.Unmarshal(reqBody, body)
				if err != nil {
					logger.Print("KeyAuth", "err", err.Error())
					return false, nil
				}
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
			}

			md5Key := utils.GenMd5Key(body.FbUserID)
			logger.Print(c.Request().URL, "KeyAuth", requestID, body.FbUserID, key, md5Key, key == md5Key)
			return key == md5Key, nil
		}))
	}
	InitApi(e)
	InitAdmin(e)

	if configs.EnvPath == "dev" {
		port = os.Getenv("PORT")
	}
	return e.Start(":" + port)
}

func Shutdown() error {
	logger.Print("showdown db ...")
	if err := mysql.Close(); err != nil {
		return err
	}
	logger.Print("showdown db done.")
	return nil
}
