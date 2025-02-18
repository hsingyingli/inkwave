package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_URL string `mapstructure:"DB_URL"`
}

func LoadEnv() (Config, error) {
	var cfg Config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
