package routes

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hero/configs"
	"hero/pkg/logger"
	"io/ioutil"
)

func Run() {
	port := configs.Get("server.port")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	if configs.EnvPath != "local" {
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			logger.Print(c.Request().URL, "BodyDump reqBody", requestID, string(reqBody))
			logger.Print(c.Request().URL, "BodyDump resBody", requestID, string(resBody))
		}))
		//安全性驗證
		e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			body := &struct {
				FbUserID string `json:"fb_user_id"`
			}{}
			if c.Request().Body != nil { // Read
				reqBody, _ := ioutil.ReadAll(c.Request().Body)
				err := json.Unmarshal(reqBody, body)
				if err != nil {
					return false, nil
				}
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
			}

			h := md5.New()
			h.Write([]byte(body.FbUserID))
			md5Key := hex.EncodeToString(h.Sum(nil))
			logger.Print(c.Request().URL, "KeyAuth", requestID, body.FbUserID, key, md5Key)
			return key == md5Key, nil
		}))
	}
	InitApi(e)
	InitAdmin(e)
	e.Logger.Fatal(e.Start(":" + port))
}
