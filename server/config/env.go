package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	DATABASE_URL string
	PORT         int
}

func NewEnv() *Env {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetDefault("PORT", 8080)

	return &Env{
		DATABASE_URL: viper.GetString("DATABASE_URL"),
		PORT:         viper.GetInt("PORT"),
	}
}
