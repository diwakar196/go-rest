package database

import (
	"fmt"
	"log"
	"time"

	"go-rest/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *gorm.DB
	var err error

	// Retry for 30 seconds (10 attempts)
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database.")
			return db
		}

		log.Printf("Attempt %d: Error connecting to database: %v", i, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("Failed to connect to database after multiple attempts: %v", err)
	return nil
}
