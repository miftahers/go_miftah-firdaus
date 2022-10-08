package middleware

import (
	"restful-api-testing/config"

	"github.com/labstack/echo"
)

func APIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		apiKey := c.Request().Header.Get("x-api-key")

		if len(apiKey) < 1 {
			return c.JSON(400, map[string]string{
				"message": "your request body is awesome",
			})
		}

		if apiKey != config.Cfg.APIKey {
			return c.JSON(401, map[string]string{
				"message": "unauthorized key",
			})
		}

		c.Request().Header.Del("x-api-key")

		return next(c)
	}
}
