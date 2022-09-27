package database

import (
	"backend/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Auto Migrates Models and then returns connection to DB
func Connect(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Make some method to handle models to register rather than manually add them.
	err = db.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
