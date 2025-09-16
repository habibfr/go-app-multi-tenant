package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"gorm.io/gorm"
)

type ProductSeed struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	TenantName string  `json:"tenant_name"`
}

func ListProductSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./database/seeders/json/products.json")
	if err != nil {
		return err
	}

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listProduct []ProductSeed
	if err := json.Unmarshal(jsonData, &listProduct); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entities.Product{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entities.Product{}); err != nil {
			return err
		}
	}

	for _, data := range listProduct {
		var tenant entities.Tenant
		if err := db.Where("name = ?", data.TenantName).First(&tenant).Error; err != nil {
			return err // tenant tidak ditemukan
		}

		var product entities.Product
		err := db.Where(&entities.Product{Name: data.Name}).First(&product).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&product, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			newProduct := entities.Product{
				Name:     data.Name,
				Price:    data.Price,
				TenantID: tenant.ID,
			}
			if err := db.Create(&newProduct).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// func ListUserSeeder(db *gorm.DB) error {
// 	jsonFile, err := os.Open("./database/seeders/json/users.json")
// 	if err != nil {
// 		return err
// 	}

// 	jsonData, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		return err
// 	}

// 	var listUser []entities.User
// 	if err := json.Unmarshal(jsonData, &listUser); err != nil {
// 		return err
// 	}

// 	hasTable := db.Migrator().HasTable(&entities.User{})
// 	if !hasTable {
// 		if err := db.Migrator().CreateTable(&entities.User{}); err != nil {
// 			return err
// 		}
// 	}

// 	for _, data := range listUser {
// 		var user entities.User
// 		err := db.Where(&entities.User{Email: data.Email}).First(&user).Error
// 		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 			return err
// 		}

// 		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
// 		if isData == 0 {
// 			if err := db.Create(&data).Error; err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
