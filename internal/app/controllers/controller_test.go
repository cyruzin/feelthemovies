package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/cyruzin/feelthemovies/internal/app/config"
	model "github.com/cyruzin/feelthemovies/internal/app/models"
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
	defer tearDownHandlers(h)
	os.Exit(m.Run())
}

func databaseConn(logger *logger.Logger) *sqlx.DB {
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost,
		cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", url)
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

func redisConn(logger *logger.Logger) *re.Client {
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	client := re.NewClient(&re.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPass,
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		logger.Fatal("Could not open connection to Redis: ", err)
	}

	logger.Info("Redis: Connection OK.")

	return client
}

func initHandlers() *setupTest {
	logger, err := logger.Init()
	if err != nil {
		panic("Could not initiate the logger: " + err.Error())
	}

	database := databaseConn(logger)
	redis := redisConn(logger)
	model := model.New(database)
	validator := validator.New()
	controllers := New(
		model,
		redis,
		validator,
		logger,
	)

	return &setupTest{
		controllers,
		database,
		redis,
		logger,
	}
}

func tearDownHandlers(h *setupTest) {
	h.database.Close()
	h.redis.Close()
	h.logger.Sync()
}

func TestSetCache(t *testing.T) {
	testKey := struct {
		Name string
	}{
		"Test",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.handler.SetCache(ctx, "testKey", &testKey); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveCache(t *testing.T) {
	testKey := struct {
		Name string
	}{
		"Test",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.handler.SetCache(ctx, "testKey", &testKey); err != nil {
		t.Fatal(err)
	}

	if err := h.handler.RemoveCache(ctx, "testKey"); err != nil {
		t.Fatal(err)
	}
}

func TestCheckCache(t *testing.T) {
	testKey := struct {
		Name string
	}{
		"Test",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.handler.SetCache(ctx, "testKey", &testKey); err != nil {
		t.Fatal(err)
	}

	cacheKey := struct{ Name string }{}

	cache, err := h.handler.CheckCache(ctx, "testKey", &cacheKey)
	if err != nil {
		t.Fatal(err)
	}

	if !cache {
		t.Errorf("Cache key differs. Expected %t.\n Got %t", true, cache)
	}
}

func TestIDParse(t *testing.T) {
	id, err := h.handler.IDParser("1")
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Errorf("ID differs. Expected %d.\n Got %d", 1, id)
	}
}

func TestPageParser(t *testing.T) {
	params := url.Values{}
	params["page"] = []string{"1"}

	id, err := h.handler.PageParser(params)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Errorf("ID differs. Expected %d.\n Got %d", 1, id)
	}
}

func TestToJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/recommendations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.handler.GetRecommendations)

	handler.ServeHTTP(rr, req)

	h.handler.ToJSON(rr, http.StatusOK, &model.Recommendation{})
}
