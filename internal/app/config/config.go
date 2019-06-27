package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config is a configuration struct that contains
// all environment variables of the app.
type Config struct {
	ServerPort   string `envconfig:"SERVERPORT"`
	DBHost       string `envconfig:"DBHOST"`
	DBUser       string `envconfig:"DBUSER"`
	DBName       string `envconfig:"DBNAME"`
	DBPort       string `envconfig:"DBPORT"`
	DBPass       string `envconfig:"DBPASS"`
	RedisAddress string `envconfig:"REDISADDR"`
	RedisPass    string `envconfig:"REDISPASS"`
	NewRelicKey  string `envconfig:"NEWRELICKEY"`
	JWTSecret    string `envconfig:"JWTSECRET"`
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
