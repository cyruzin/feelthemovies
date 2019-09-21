package handler

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cyruzin/feelthemovies/internal/app/config"
	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"
	"github.com/go-chi/chi"
	re "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	h      = initHandlers()
	router = chi.NewRouter()
	v      *validator.Validate
	info   = &model.Auth{
		ID:    1,
		Name:  "Admin",
		Email: "admin@admin.com",
	}
)

type setupTest struct {
	handler  *Setup
	database *sqlx.DB
	redis    *re.Client
	logger   *logger.Logger
}

func TestMain(m *testing.M) {
	router.Use(h.handler.JSONMiddleware)
	defer tearDownHandlers(h) // Closing connections.
	os.Exit(m.Run())
}

func databaseConn() *sqlx.DB {
	cfg, err := config.Load() // Loading environment variables.
	if err != nil {
		log.Fatal(err.Error())
	}

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		log.Fatal("Could not open connection to MySQL: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to MySQL: ", err)
	}

	log.Println("MySQL: Connection OK.")
	return db
}

func redisConn() *re.Client {
	cfg, err := config.Load() // Loading environment variables.
	if err != nil {
		log.Fatal(err.Error())
	}

	client := re.NewClient(&re.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPass,
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		log.Fatal("Could not open connection to Redis: ", err)
	}

	log.Println("Redis: Connection OK.")

	return client
}

func initHandlers() *setupTest {
	loggerInstance, err := logger.Init() // Uber Zap Logger instance.
	if err != nil {
		loggerInstance.Fatal("Could not initiate the logger: " + err.Error())
	}

	databaseInstance := databaseConn()               // Database instance.
	redisInstance := redisConn()                     // Redis client instance.
	modelInstance := model.Connect(databaseInstance) // Passing database instance to the model pkg.
	validatorInstance := validator.New()             // Validator instance.
	handlersInstance := NewHandler(
		modelInstance,
		redisInstance,
		validatorInstance,
		loggerInstance,
	) // Passing instances to the handlers pkg.

	return &setupTest{
		handlersInstance,
		databaseInstance,
		redisInstance,
		loggerInstance,
	}
}

func tearDownHandlers(h *setupTest) {
	h.database.Close()
	h.redis.Close()
	h.logger.Sync()
}
