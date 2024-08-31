package userrepository

import (
	"context"

	"org.idev.bunny/backend/component/redis"
	userentity "org.idev.bunny/backend/domain/user/entity"
)

type UserCacheRepo struct {
	redisCli *redis.Client
}

// create new instance of redis repository
func NewCacheRepository(redisCli *redis.Client) *UserCacheRepo {
	return &UserCacheRepo{redisCli: redisCli}
}

// Generate by command:
// GoImpl org.idev.bunny/backend/domain/user.UserCacheRepository
func (c *UserCacheRepo) FindById(ctx context.Context, userId string) (*userentity.User, error) {
	var user userentity.User
	err := c.redisCli.GetStruct(ctx, userId, &user)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (c *UserCacheRepo) Save(ctx context.Context, user userentity.User) error {
	return c.redisCli.Set(ctx, user.UserId, user, 0)
}
