package infrastructure

import (
	"app/internal/domain"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Todo{})
	if err != nil {
		log.Fatal("DB migration failed: ", err)
	}
	log.Println("Migration completed")
}
