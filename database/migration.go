package database

import (
	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entities.Tenant{},
		&entities.User{},
		&entities.RefreshToken{},
		&entities.Product{},
	); err != nil {
		return err
	}

	return nil
}

func MigrateFresh(db *gorm.DB) error {
	// Drop tables
	if err := db.Migrator().DropTable(
		&entities.Tenant{},
		&entities.User{},
		&entities.RefreshToken{},
		&entities.Product{},
	); err != nil {
		return err
	}
	// Migrate ulang
	return Migrate(db)
}
