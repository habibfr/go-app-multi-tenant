package database

import (
	"github.com/Caknoooo/go-gin-clean-starter/database/seeders/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		seeds.ListUserSeeder,
		seeds.ListTenantSeeder,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
