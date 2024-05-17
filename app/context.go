package app

import (
	"context"

	"org.idev.bunny/backend/component/kafka"
	"org.idev.bunny/backend/component/redis"
	sqlc_generated "org.idev.bunny/backend/generated/sqlc"
)

// Contain app context
type AppContext struct {
	Ctx           context.Context
	Config        *appConfig
	Db            sqlc_generated.DBTX
	RedisCli      *redis.RedisClient
	KafkaProducer *kafka.Producer
}
