package userdomain

import (
	"context"

	userentity "org.idev.bunny/backend/domain/user/entity"
)

type UserRepository interface {
	FindByUserId(ctx context.Context, userId string) (*userentity.User, error)
	Save(ctx context.Context, user *userentity.User) (*userentity.User, error)
}

type UserUseCase interface {
	FindByUserId(ctx context.Context, userId string) (*userentity.User, error)
	Create(ctx context.Context, userId string, username string) (*userentity.User, error)
}
