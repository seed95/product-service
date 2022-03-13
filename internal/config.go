package internal

import (
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	PostgresUri string
}

type Config struct {
	ProductRepo PostgresConfig
}

func NewConfig(prefix string) *Config {
	// Viper Configuration
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return &Config{
		ProductRepo: PostgresConfig{
			PostgresUri: v.GetString("postgres_dsn"),
		},
	}
}
