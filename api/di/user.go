package di

import (
	"org.idev.bunny/backend/component/redis"
	userdomain "org.idev.bunny/backend/domain/user"
	sqlc_generated "org.idev.bunny/backend/generated/sqlc"
	userrepository "org.idev.bunny/backend/repository/user"
)

func NewUserUseCase(queries *sqlc_generated.Queries, redisCli *redis.RedisClient) userdomain.UserUseCase {
	repo := NewUserRepository(queries, redisCli)
	return userdomain.NewUserUseCase(repo)
}

func NewUserRepository(queries *sqlc_generated.Queries, redisCli *redis.RedisClient) userdomain.UserRepository {
	sql := userrepository.NewSqlRepository(queries)
	cache := userrepository.NewCacheRepository(redisCli)
	return userrepository.New(sql, cache)
}
