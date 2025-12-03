package migrations

import (
	"crud-service/domain/entities"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{}, // add models here
		// &entities.Product{},
		// &entities.Order{},
	)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
