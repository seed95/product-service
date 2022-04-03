package internal

import (
	"github.com/spf13/viper"
	"time"
)

type LogConfig struct {
	StdLevel int // Standard core log level
}

type PostgresConfig struct {
	DSN string
}

type Config struct {
	Log            LogConfig
	ServiceTimeout time.Duration
	ProductRepo    PostgresConfig
	GRPCPort       string
}

func NewConfig(prefix string) *Config {
	// Viper Configuration
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	return &Config{
		Log: LogConfig{
			StdLevel: v.GetInt("std_level"),
		},
		ServiceTimeout: v.GetDuration("timeout"),
		ProductRepo: PostgresConfig{
			DSN: v.GetString("postgres_dsn"),
		},
		GRPCPort: v.GetString("grpc_port"),
	}
}
