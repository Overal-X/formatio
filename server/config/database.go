package config

import (
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(i *do.Injector) (*gorm.DB, error) {
	env := do.MustInvoke[*Env](i)

	return gorm.Open(postgres.Open(env.DATABASE_URL), &gorm.Config{})
}
