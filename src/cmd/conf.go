package main

import (
	"os"

	"belajar-go/src/config/database"
	"belajar-go/src/config/logger"
	"belajar-go/src/config/server"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   server.Config     `mapstructure:"server"`
	Gin      server.GinConfig  `mapstructure:"gin"`
	Logger   logger.Config     `mapstructure:"logger"`
	Postgres database.DBConfig `mapstructure:"postgres"`
}

func LoadConfig() (*AppConfig, error) {

	viper.SetConfigFile("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg AppConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.User = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.DBName = os.Getenv("DB_NAME")
	cfg.Postgres.SSLMode = os.Getenv("DB_SSLMODE")

	return &cfg, nil
}
