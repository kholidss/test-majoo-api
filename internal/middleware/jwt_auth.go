package middleware

import (
	"net/http"
	"strings"

	"test-majoo-api/internal/utils"
	"test-majoo-api/internal/utils/token"

	"github.com/labstack/echo/v4"
)

func JwtAuth() echo.MiddlewareFunc {
	return jwtAuthMiddleware
}

func jwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// cfg := config.LoadConfig()

		jwtMaker := c.Get("token-maker").(token.Maker)

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)

		payload, err := jwtMaker.VerifyToken(idToken)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.Unauthorized(err.Error()))
		}

		c.Set("admin", payload)

		return next(c)
	}

}
