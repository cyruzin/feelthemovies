package logger

import (
	"github.com/cyruzin/feelthemovies/internal/app/config"
	"go.uber.org/zap"
)

// Logger is a instance of Uber's zap logger.
type Logger struct {
	*zap.SugaredLogger
}

// Init initiates the logger.
func Init() (*Logger, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log, err := zap.NewDevelopment()

	if env := cfg.EnvMode; env == "production" {
		log, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{log.Sugar()}, nil
}
