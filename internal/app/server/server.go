package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/InVisionApp/go-health/handlers"
	"github.com/jmoiron/sqlx"

	health "github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/checkers"
	redisCheck "github.com/InVisionApp/go-health/checkers/redis"
	"github.com/cyruzin/feelthemovies/internal/app/config"
	"github.com/cyruzin/feelthemovies/internal/app/controllers"
	"github.com/cyruzin/feelthemovies/internal/app/models"
	"github.com/cyruzin/feelthemovies/internal/app/router"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"

	re "github.com/go-redis/redis"
	validator "gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

// New initiates the server.
func New(ctx context.Context) {
	cfg, err := config.Load()
	if err != nil {
		panic(err.Error())
	}

	logger, err := logger.Init()
	if err != nil {
		panic("Could not initiate the logger: " + err.Error())
	}

	database := database(ctx, cfg, logger)
	redis := redis(ctx, cfg, logger)
	model := models.New(database)
	validator := validator.New()
	controllers := controllers.New(
		model,
		redis,
		validator,
		logger,
	)

	defer logger.Sync()
	defer database.Close()
	defer redis.Close()

	healthCheck, err := healthChecks(cfg, database)
	if err != nil {
		logger.Info("Failed to perform health checks")
	}

	r := router.New(
		controllers,
		handlers.NewJSONHandlerFunc(healthCheck, nil),
		logger,
	)

	srv := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		ReadTimeout:       cfg.ReadTimeOut,
		ReadHeaderTimeout: cfg.ReadHeaderTimeOut,
		WriteTimeout:      cfg.WriteTimeOut,
		IdleTimeout:       cfg.IdleTimeOut,
		Handler:           r,
	}

	// Graceful shutdown setup.
	idleConnsClosed := make(chan struct{})

	go func() {
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, os.Interrupt)
		<-gracefulStop

		logger.Info("Shutting down the server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Errorf("Server failed to shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	logger.Info("Listening on port: " + cfg.ServerPort)
	logger.Info("You're good to go! :)")

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Errorf("Server failed to start: %v", err)
	}
	<-idleConnsClosed
}

// Database connection.
func database(
	ctx context.Context,
	cfg *config.Config,
	logger *logger.Logger,
) *sqlx.DB {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.ConnectContext(ctx, "mysql", url)
	if err != nil {
		logger.Fatal("Could not open connection to MySQL: ", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("Could not connect to MySQL: ", err)
	}

	logger.Info("MySQL: Connection OK.")
	return db
}

// Redis connection.
func redis(
	ctx context.Context,
	cfg *config.Config,
	logger *logger.Logger,
) *re.Client {
	client := re.NewClient(&re.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPass,
		DB:       0,
	})

	_, err := client.WithContext(ctx).Ping().Result()
	if err != nil {
		logger.Fatal("Could not open connection to Redis: ", err)
	}

	logger.Info("Redis: Connection OK.")

	return client
}

// healthChecks checks the services health periodically.
func healthChecks(
	cfg *config.Config,
	db *sqlx.DB,
) (*health.Health, error) {
	h := health.New()
	h.DisableLogging()

	mysqlDB, err := checkers.NewSQL(&checkers.SQLConfig{
		Pinger: db,
	})
	if err != nil {
		return nil, err
	}

	redisDB, err := redisCheck.NewRedis(
		&redisCheck.RedisConfig{
			Auth: &redisCheck.RedisAuthConfig{
				Addr:     cfg.RedisAddress,
				Password: cfg.RedisPass,
				DB:       0,
			},
			Ping: true,
		})
	if err != nil {
		return nil, err
	}

	if err = h.AddChecks([]*health.Config{
		{
			Name:     "feelthemovies-database",
			Checker:  mysqlDB,
			Interval: time.Duration(3) * time.Second,
			Fatal:    true,
		},
		{
			Name:     "feelthemovies-redis",
			Checker:  redisDB,
			Interval: time.Duration(3) * time.Second,
			Fatal:    true,
		},
	}); err != nil {
		return nil, err
	}

	if err := h.Start(); err != nil {
		return nil, err
	}

	return h, nil
}
