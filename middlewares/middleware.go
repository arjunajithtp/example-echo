package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// ValidateToken is the middleware to validate the test token
func ValidateToken(key string, c echo.Context) (bool, error) {
	return key == "test-token", nil
}

// BasicAuthValidator is the middleware to validate the test basic auth
func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	if username == "test-name" && password == "password" {
		return true, nil
	}
	return false, nil
}

// CustomAuth middleware is for custom authentication
func CustomAuth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		tokenPieces := strings.Split(token, " ")
		if len(tokenPieces) == 2 && tokenPieces[0] == "Basic" {
			// TODO: do basic auth
			return handler(c)
		}
		// TODO: do session auth

		// if both validations fails
		return c.JSON(http.StatusUnauthorized, "error while trying to authenticate")
	}
}
