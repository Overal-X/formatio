package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(env *Env) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(env.DATABASE_URL), &gorm.Config{})
}
