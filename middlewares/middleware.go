package middlewares

import (
	"github.com/labstack/echo/v4"
)

// ValidateToken is the middleware to validate the test token
func ValidateToken(key string, c echo.Context) (bool, error) {
	return key == "test-token", nil
}
