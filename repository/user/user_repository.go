package userrepository

import (
	"context"

	"org.idev.bunny/backend/common/logger"
	userentity "org.idev.bunny/backend/domain/user/entity"
	usercache "org.idev.bunny/backend/repository/user/cache"
	usersql "org.idev.bunny/backend/repository/user/sql"
)

type UserRepository struct {
	sqlRepo   *usersql.UserSqlRepo
	cacheRepo *usercache.UserCacheRepo
}

func New(sqlRepo *usersql.UserSqlRepo, cacheRepo *usercache.UserCacheRepo) *UserRepository {
	return &UserRepository{
		sqlRepo:   sqlRepo,
		cacheRepo: cacheRepo,
	}
}

func (r *UserRepository) FindByUserId(ctx context.Context, userId string) (*userentity.User, error) {

	log := logger.New("UserRepository", "FindByUserId")

	if user, err := r.cacheRepo.FindById(ctx, userId); err == nil {
		return user, nil
	}
	user, err := r.sqlRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	if err := r.cacheRepo.Save(ctx, *user); err != nil {
		log.Errorf("failed to save user to cache: %v", err)
	}
	return user, nil
}

func (r *UserRepository) Save(ctx context.Context, u *userentity.User) (*userentity.User, error) {
	return r.sqlRepo.Save(ctx, u)
}
