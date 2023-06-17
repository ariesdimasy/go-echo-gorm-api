package config

import (
	// "fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseRepository struct {
	DB *gorm.DB
}

func Connection() (*gorm.DB, error) {
	// Connect to the database
	dsn := "root@tcp(localhost:3306)/db_belajar_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// func Migration() {
// 	// Migrate the database
// 	db, err := Connection().db.AutoMigrate(&models.Product{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
