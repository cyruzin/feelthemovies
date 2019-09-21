package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is a configuration struct that contains
// all environment variables of the app.
type Config struct {
	EnvMode      string `envconfig:"ENVMODE" required:"true" default:"development"`
	ServerPort   string `envconfig:"SERVERPORT" required:"true" default:"8000"`
	DBHost       string `envconfig:"DBHOST" required:"true" default:"localhost"`
	DBPort       string `envconfig:"DBPORT" required:"true" default:"3306"`
	DBUser       string `envconfig:"DBUSER" required:"true" default:"root"`
	DBName       string `envconfig:"DBNAME" required:"true" default:"api_feelthemovies"`
	DBPass       string `envconfig:"DBPASS" required:"true" default:"secret"`
	RedisAddress string `envconfig:"REDISADDR" required:"true" default:"localhost:6379"`
	RedisPass    string `envconfig:"REDISPASS" required:"true" default:"secret"`
	NewRelicKey  string `envconfig:"NEWRELICKEY"`
	JWTSecret    string `envconfig:"JWTSECRET" required:"true" default:"secret"`
}

// Load loads the app the configuration based
// in the environment variables.
func Load() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
