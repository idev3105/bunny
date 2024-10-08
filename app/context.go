package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"org.idev.bunny/backend/component/kafka"
	"org.idev.bunny/backend/component/mongo"
	"org.idev.bunny/backend/component/redis"
)

// Contain app context
type AppContext struct {
	Ctx           context.Context
	Config        *Config
	Db            *pgxpool.Pool
	Redis         *redis.Client
	KafkaProducer *kafka.Producer
	Mongo         *mongo.Client
}
