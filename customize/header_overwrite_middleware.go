package customize

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func overwrite(req http.Request, res echo.Response) {
	res.Header().Set(echo.HeaderAccessControlAllowOrigin, req.Header.Get(echo.HeaderOrigin))
}

func HeaderOverwriteMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Before(func() {
			overwrite(*ctx.Request(), *ctx.Response())
		})

		return next(ctx)
	}
}
