package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"org.idev.bunny/backend/api/di"
	"org.idev.bunny/backend/app"
	sqlc_generated "org.idev.bunny/backend/generated/sqlc"
	tokenutil "org.idev.bunny/backend/utils/token"
)

// User gin handler
type UserHandler struct {
	appCtx *app.AppContext
}

func NewUserHandler(appCtx *app.AppContext) *UserHandler {
	return &UserHandler{appCtx: appCtx}
}

// @Id CreateUser
// @Summary Create new user
// @Description Create new user
// @Tags user
// @Accept json
// @Produce json
// @Param idToken body string true "idToken"
// @Param username body string true "username"
// @Success 200 {object} UserDto
// @Router /user [post]
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

		userUseCase := di.NewUserUseCase(sqlc_generated.New(u.appCtx.Db), u.appCtx.Redis)
		user, err := userUseCase.Create(ctx.Request().Context(), userId, data.Username)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}

// @Id GetUserByUserId
// @Summary Get user by userId
// @Description Get user by userId
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "userId"
// @Success 200
// @Router /user/{id} [get]
func (u *UserHandler) GetUserByUserId() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("id")

		userUseCase := di.NewUserUseCase(sqlc_generated.New(u.appCtx.Db), u.appCtx.Redis)
		user, err := userUseCase.FindByUserId(ctx.Request().Context(), userId)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}
