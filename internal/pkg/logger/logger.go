package logger

import (
	"go.uber.org/zap"
)

// Logger is a instance of Uber's zap logger.
type Logger struct {
	*zap.SugaredLogger
}

// Init initiates the logger.
func Init() (*Logger, error) {
	log, err := zap.NewProduction() // Uber Zap Logger instance.
	if err != nil {
		return nil, err
	}
	return &Logger{log.Sugar()}, nil
}
