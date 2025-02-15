package config

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[Request]: ${time_custom} ${remote_ip} ${method} ${uri} ${status} - ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	return e
}
