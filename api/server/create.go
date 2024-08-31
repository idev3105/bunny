package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	appMiddleware "org.idev.bunny/backend/api/middleware"
	"org.idev.bunny/backend/api/route"
	"org.idev.bunny/backend/app"
	"org.idev.bunny/backend/common/errors"
	"org.idev.bunny/backend/common/logger"
	"org.idev.bunny/backend/component/kafka"
	"org.idev.bunny/backend/component/mongo"
	"org.idev.bunny/backend/component/redis"
)

func Create(ctx context.Context) (*Server, error) {
	log := logger.New("Server", "create server")

	appConfig, err := app.LoadConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	e := setupEcho()

	AppCtx = &app.AppContext{
		Ctx:    ctx,
		Config: appConfig,
	}

	if err := setupComponents(ctx, AppCtx, appConfig, log); err != nil {
		return nil, errors.Wrap(err, "failed to setup components")
	}

	e.Use(appMiddleware.AuthGuard(AppCtx))

	setupRoutes(e, AppCtx)

	return &Server{e: e}, nil
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}

func setupComponents(ctx context.Context, appCtx *app.AppContext, config *app.Config, log *logger.Logger) error {
	var err error

	if config.EnableRedis {
		appCtx.Redis, err = setupRedis(ctx, config.RedisUrl, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup Redis")
		}
	}

	if config.EnableDb {
		appCtx.Db, err = setupDatabase(ctx, config.DbUrl, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup database")
		}
	}

	if config.EnableMongo {
		appCtx.Mongo, err = setupMongo(ctx, config.MongoUrl, config.MongoDbName, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup MongoDB")
		}
	}

	if config.EnableKafka {
		appCtx.KafkaProducer, err = setupKafka(config.KafkaHost, config.KafkaPort, log)
		if err != nil {
			return errors.Wrap(err, "failed to setup Kafka")
		}
	}

	return nil
}

func setupRedis(ctx context.Context, url string, log *logger.Logger) (*redis.Client, error) {
	log.Info("Connecting to Redis: " + url)
	return redis.NewClient(ctx, url)
}

func setupDatabase(ctx context.Context, url string, log *logger.Logger) (*pgxpool.Pool, error) {
	log.Info("Connecting to database: " + url)
	poolCfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse database config")
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create database pool")
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}
	return pool, nil
}

func setupMongo(ctx context.Context, url, dbName string, log *logger.Logger) (*mongo.Client, error) {
	log.Info("Connecting to MongoDB: " + url)
	return mongo.NewMongoClient(ctx, url, dbName)
}

func setupKafka(host string, port int32, log *logger.Logger) (*kafka.Producer, error) {
	log.Info(fmt.Sprintf("Connecting to Kafka cluster: %s:%d", host, port))
	return kafka.NewProducer(host, port)
}

func setupRoutes(e *echo.Echo, appCtx *app.AppContext) {
	route.NewExamplePanicErrorRouter(e)

	v1 := e.Group("/api/v1")
	route.NewUserRouter(v1, appCtx)
}
