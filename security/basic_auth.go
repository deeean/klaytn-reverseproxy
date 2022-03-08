package security

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth(username, password string) middleware.BasicAuthValidator {
	return func(u string, p string, context echo.Context) (bool, error) {
		if u == username && p == password {
			return true, nil
		}

		return false, nil
	}
}
