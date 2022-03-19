package internal

import (
	"github.com/spf13/viper"
	"time"
)

type PostgresConfig struct {
	PostgresUri string
}

type Config struct {
	ServiceTimeout time.Duration
	ProductRepo    PostgresConfig
}

func NewConfig(prefix string) *Config {
	// Viper Configuration
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return &Config{
		ServiceTimeout: v.GetDuration("timeout"),
		ProductRepo: PostgresConfig{
			PostgresUri: v.GetString("postgres_dsn"),
		},
	}
}
