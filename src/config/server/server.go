package server

import "time"

type Config struct {
	Port            string        `mapstructure:"port"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	Mode            string        `mapstructure:"mode"`
}

type GinConfig struct {
	AppName string `mapstructure:"app_name"`
}
