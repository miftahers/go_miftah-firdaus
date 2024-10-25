package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip},id=${id},bytes_in=${bytes_in}, bytes_out=${bytes_out}, latency_human=${latency_human},\n",
	}))
}
