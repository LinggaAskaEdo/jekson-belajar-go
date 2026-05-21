package database

import (
	"fmt"
	"time"
)

type DBConfig struct {
	Enabled         bool          `mapstructure:"enabled"`
	Driver          string        `mapstructure:"driver"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
