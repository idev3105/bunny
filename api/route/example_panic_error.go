package route

import "github.com/labstack/echo/v4"

func NewExamplePanicErrorRouter(e *echo.Echo) {
	e.GET("/example-panic-error", func(ctx echo.Context) error {
		panic("example panic error")
	})
}
