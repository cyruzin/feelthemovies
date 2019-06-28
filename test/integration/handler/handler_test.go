package handler_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"

	"github.com/cyruzin/feelthemovies/internal/app/handler"
	"github.com/cyruzin/feelthemovies/internal/pkg/logger"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/test/integration/setup"
	"gopkg.in/go-playground/validator.v9"
)

var (
	h    = initHandlers()
	r    = chi.NewRouter()
	v    *validator.Validate
	info = &model.Auth{
		ID:    600,
		Name:  "Admin",
		Email: "admin@admin.com",
	}
)

func TestMain(m *testing.M) {
	r.Use(h.h.JSONMiddleware)
	defer tearDownHandlers(h) // Closing connections.
	os.Exit(m.Run())
}

type setupTest struct {
	h  *handler.Setup
	db *sql.DB
	rc *redis.Client
	l  *logger.Logger
}

func initHandlers() *setupTest {
	l, err := logger.Init() // Uber Zap Logger instance.
	if err != nil {
		l.Fatal("Could not initiate the logger: " + err.Error())
	}

	db := setup.Database()                // Database instance.
	rc := setup.Redis()                   // Redis client instance.
	mc := model.Connect(db)               // Passing database instance to the model pkg.
	v = validator.New()                   // Validator instance.
	h := handler.NewHandler(mc, rc, v, l) // Passing instances to the handlers pkg.

	return &setupTest{
		h, db, rc, l,
	}
}

func tearDownHandlers(h *setupTest) {
	h.db.Close()
	h.rc.Close()
	h.l.Sync()
}
