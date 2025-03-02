package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_URL                   string        `mapstructure:"DB_URL"`
	ACCESS_TOKEN_SECRET_KEY  string        `mapstructure:"ACCESS_TOKEN_SECRET_KEY"`
	ACCESS_TOKEN_DURATION    time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REFRESH_TOKEN_SECRET_KEY string        `mapstructure:"REFRESH_TOKEN_SECRET_KEY"`
	REFRESH_TOKEN_DURATION   time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadEnv() (*Config, error) {
	cfg := new(Config)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
