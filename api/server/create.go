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
	"org.idev.bunny/backend/common/logger"
	"org.idev.bunny/backend/component/kafka"
	"org.idev.bunny/backend/component/redis"
	userdomain "org.idev.bunny/backend/domain/user"
	userrepository "org.idev.bunny/backend/repository/user"
	usercache "org.idev.bunny/backend/repository/user/cache"
	usersql "org.idev.bunny/backend/repository/user/sql"
)

func Create() *Server {

	ctx := context.Background()

	e := echo.New()
	log := logger.New("Server", "create server")

	e.Debug = true

	e.Use(middleware.Recover()) // recover from panic, use can custom error handler
	e.Use(middleware.Logger())  // show http request logs

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appConfig := app.LoadConfig()

	log.Info("Connect to redis " + appConfig.RedisUrl)
	redisCli, err := redis.NewClient(appConfig.RedisUrl)
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	log.Info("Connect to database " + appConfig.Dsn)
	poolCfg, err := pgxpool.ParseConfig(appConfig.Dsn)
	if err != nil {
		panic(err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		panic(err)
	}

	log.Info("Connect to kafka cluster " + appConfig.KafkaHost + ":" + fmt.Sprint(appConfig.KafkaPort))
	kafkaProducer, err := kafka.NewProducer(appConfig.KafkaHost, appConfig.KafkaPort)
	if err != nil {
		log.Fatalf("Failed to connect to kafka: %v", err)
	}

	// init app context instance
	AppCtx = &app.AppContext{
		Ctx:           ctx,
		Config:        appConfig,
		Db:            pool,
		RedisCli:      redisCli,
		KafkaProducer: kafkaProducer,
	}

	userRepo := userrepository.New(usersql.NewSqlRepository(AppCtx.Db), usercache.NewCachRepository(AppCtx.RedisCli))
	userUsecase := userdomain.NewUserUseCase(userRepo)

	e.Use(appMiddleware.AuthGuard(AppCtx, userRepo))

	route.NewExamplePanicErrorRouter(e)

	v1 := e.Group("/api/v1")
	{
		route.NewUserRouter(v1, AppCtx, userUsecase)
	}

	return &Server{e: e}
}
