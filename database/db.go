package database

import (
	"fmt"
	"log"
	"finalproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host       = "localhost"
	user       = "postgres"
	password   = "admin"
	dbPort     = 5432
	dbname     = "finalproject"
	debug_mode = true
	db         *gorm.DB
	err        error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s dbPort=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	fmt.Println("Sukses koneksi ke database")
	if debug_mode == true {
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

}

func GetDB() *gorm.DB {
	return db
}
