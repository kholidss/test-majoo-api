package middleware

import (
	"context"
	"net/http"
	"strings"

	"test-majoo-api/internal/utils"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

func FirebaseAuth() echo.MiddlewareFunc {
	return firebaseAuthMiddleware
}

func firebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		client := c.Get("firebaseAuth").(*auth.Client)
		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, utils.Unauthorized("Authorization Failed"))
		}

		c.Set("user", token)
		return next(c)
	}
}
