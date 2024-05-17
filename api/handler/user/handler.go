package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"org.idev.bunny/backend/app"
	tokenutil "org.idev.bunny/backend/common/util/token"
	userdomain "org.idev.bunny/backend/domain/user"
)

// User gin handler
type UserHandler struct {
	appCtx      *app.AppContext
	userUseCase userdomain.UserUseCase
}

func NewUserHandler(appCtx *app.AppContext, userUserCase userdomain.UserUseCase) *UserHandler {
	return &UserHandler{appCtx: appCtx, userUseCase: userUserCase}
}

// Create new user
// @Summary Create new user
// @Description Create new user
// @Tags user
// @Accept json
// @Produce json
// @Param idToken body string true "idToken"
// @Param username body string true "username"
// @Success 200 {object} UserDto
// @Router /user [post]
// @Security ApiKeyAuth
func (u *UserHandler) CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// bind body from request
		var data CreateUserRequest
		err := (&echo.DefaultBinder{}).BindBody(ctx, &data)
		if err != nil {
			panic(err)
		}

		token, err := tokenutil.Parse(ctx.Request().Context(), data.IdToken, u.appCtx.Config.JWKsUrl)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Token is invalid")
		}
		userId := token.Subject()

		user, err := u.userUseCase.Create(ctx.Request().Context(), userId, data.Username)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}

// Get user by user id
func (u *UserHandler) GetUserByUserId() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("id")
		user, err := u.userUseCase.FindByUserId(ctx.Request().Context(), userId)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}
