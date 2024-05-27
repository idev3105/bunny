package route

import (
	"github.com/labstack/echo/v4"
	userhandler "org.idev.bunny/backend/api/handler/user"
	"org.idev.bunny/backend/app"
)

func NewUserRouter(e *echo.Group, appCtx *app.AppContext) {
	handler := userhandler.NewUserHandler(appCtx)
	g := e.Group("/users")
	{
		g.POST("", handler.CreateUser())
		g.GET("/:id", handler.GetUserByUserId())
	}
}
