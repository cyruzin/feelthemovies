package handler_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"

	"github.com/cyruzin/feelthemovies/internal/app/handler"

	"github.com/cyruzin/feelthemovies/internal/app/model"
	"github.com/cyruzin/feelthemovies/test/integration/setup"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
)

var (
	h = initHandlers()
	r = chi.NewRouter()
	v *validator.Validate
)

func TestMain(m *testing.M) {
	defer tearDownHandlers(h)
	os.Exit(m.Run())
}

type setupTest struct {
	h  *handler.Setup
	db *sql.DB
	rc *redis.Client
	l  *zap.SugaredLogger
}

func initHandlers() *setupTest {
	l, err := zap.NewDevelopment() // Uber Zap Logger instance.

	if err != nil {
		log.Fatal("Could not initiate the logger")
	}

	db := setup.Database()                        // Database instance.
	rc := setup.Redis()                           // Redis client instance.
	mc := model.Connect(db)                       // Passing database instance to the model pkg.
	v = validator.New()                           // Validator instance.
	h := handler.NewHandler(mc, rc, v, l.Sugar()) // Passing instances to the handlers pkg.

	return &setupTest{
		h, db, rc, l.Sugar(),
	}
}

func tearDownHandlers(h *setupTest) {
	h.db.Close()
	h.rc.Close()
	h.l.Sync()
}
