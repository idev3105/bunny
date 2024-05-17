package errorhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ExampleErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)
			if err != nil {
				ctx.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
			return nil
		}
	}
}
