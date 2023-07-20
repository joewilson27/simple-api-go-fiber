package db

/*
The book model is done. Now, we configure GORM and auto migrate the model we just created.
This AutoMigrate function will create the books table for us as soon as we run this application.
*/

import (
	"log"

	"simple-go-fiber/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err!= nil {
		log.Fatalln(err)
	}

	// Migrate the model to create the "book" table
	db.AutoMigrate(&models.Book{})

	return db
}