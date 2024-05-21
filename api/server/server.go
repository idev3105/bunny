package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"org.idev.bunny/backend/app"
	"org.idev.bunny/backend/common/logger"
)

var AppCtx *app.AppContext

type Server struct {
	e *echo.Echo
}

func (s *Server) Start() error {

	log := logger.New("Server", "Start server")

	ctx, stop := signal.NotifyContext(AppCtx.Ctx, os.Interrupt, os.Kill)
	defer stop()

	go func() {
		if err := s.e.Start(fmt.Sprintf(":%v", AppCtx.Config.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(AppCtx.Ctx, 10*time.Second)
	defer cancel()

	log.Info("Close redis connection")
	if err := AppCtx.RedisCli.Close(); err != nil {
		log.Errorf("Redis close fail: %v", err)
		return err
	}

	log.Info("Close database connection")
	AppCtx.Db.(*pgxpool.Pool).Close()

	log.Info("Close mongo connection")
	if err := AppCtx.MongoClient.Close(ctx); err != nil {
		log.Errorf("Mongo close fail: %v", err)
		return err
	}

	log.Info("Close kafka producer")
	if err := AppCtx.KafkaProducer.Close(); err != nil {
		log.Errorf("Kafka producer close fail: %v", err)
		return err
	}

	if err := s.e.Shutdown(ctx); err != nil {
		log.Errorf("Server shutdown fail: %v", err)
		return err
	}
	return nil
}
