package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const header = "API_KEY"

func Auth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(header)
			if token == "" || secret != token {
				return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
			}

			return next(c)
		}
	}
}
