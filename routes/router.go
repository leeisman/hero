package routes

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hero/configs"
	"hero/pkg/logger"
	"hero/utils"
	"io/ioutil"
	"net/http"
	"os"
)

func Run() {
	port := configs.Get("server.port")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check\
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			logger.Print("echo error", requestID)
			// For invalid credentials
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			// return next(c)
		}
	})

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

	if configs.EnvPath != "local" {
		port = os.Getenv("PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
