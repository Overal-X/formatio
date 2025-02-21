package config

import (
	"fmt"

	"github.com/samber/do"
	"github.com/spf13/viper"
)

type Env struct {
	PORT         int
	DATABASE_URL string
	RABBITMQ_URL string
}

func NewEnv(i *do.Injector) (*Env, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	viper.SetDefault("PORT", 8080)

	return &Env{
		DATABASE_URL: viper.GetString("DATABASE_URL"),
		PORT:         viper.GetInt("PORT"),
		RABBITMQ_URL: viper.GetString("RABBITMQ_URL"),
	}, nil
}
