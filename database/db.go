package database

import (
	"Project-Akhir/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host         = "localhost"
	user         = "postgres"
	password     = "admin"
	databasePort = "5432"
	databaseName = "final"
)

func ReadDB() *gorm.DB {

	fix := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, databasePort)

	db, err := gorm.Open(postgres.Open(fix), &gorm.Config{})

	if err != nil {
		log.Fatal("gagal tersambung kedatabase:", err)
	}

	fmt.Println("berhasil tersambung database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	return db
}
