package scheduler

import (
	"log"
	"time"

	"github.com/Caknoooo/go-gin-clean-starter/database/entities"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	c := cron.New()
	// Setiap hari jam 00:00
	_, err := c.AddFunc("0 0 * * *", func() {
		if err := db.Model(&entities.User{}).Where("is_verified = ?", true).Update("is_verified", false).Error; err != nil {
			log.Println("Failed to update user verification:", err)
		} else {
			log.Println("All users set is_verified to false at", time.Now())
		}
	})
	if err != nil {
		log.Println("Failed to schedule cron job:", err)
	}
	c.Start()
}
