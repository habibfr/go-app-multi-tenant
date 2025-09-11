package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

func ListTenantSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/tenants.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listTenant []entities.Tenant
	if err := json.Unmarshal(jsonData, &listTenant); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entities.Tenant{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.Tenant{}); err != nil {
			return err
		}
	}

	for _, data := range listTenant {
		var tenant entities.Tenant
		err := db.Where(&entities.Tenant{Name: data.Name}).First(&tenant).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&tenant, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
