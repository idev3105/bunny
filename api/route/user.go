package route

import (
	"github.com/labstack/echo/v4"
	userhandler "org.idev.bunny/backend/api/handler/user"
	"org.idev.bunny/backend/app"
	userdomain "org.idev.bunny/backend/domain/user"
)

func NewUserRouter(e *echo.Group, appCtx *app.AppContext, userUsecase userdomain.UserUseCase) {
	handler := userhandler.NewUserHandler(appCtx, userUsecase)
	g := e.Group("/users")
	{
		g.POST("", handler.CreateUser())
		g.GET("/:id", handler.GetUserByUserId())
	}
}
