package logger

import (
	"os"

	"go.uber.org/zap"
)

// Logger is a instance of Uber's zap logger.
type Logger struct {
	*zap.SugaredLogger
}

// Init initiates the logger.
func Init() (*Logger, error) {
	log, err := zap.NewDevelopment()

	if env := os.Getenv("ENVMODE"); env == "production" {
		log, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{log.Sugar()}, nil
}
