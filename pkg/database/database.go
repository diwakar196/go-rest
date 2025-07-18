package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"go-rest/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func loadDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *gorm.DB
	var err error

	// Retry for 30 seconds (10 attempts)
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database.")
			dbConn = db
			return
		}

		log.Printf("Attempt %d: Error connecting to database: %v", i, err)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("Failed to connect to database after multiple attempts: %v", err)
}

func GetDB() *gorm.DB {
	if dbConn == nil {
		once.Do(func() {
			loadDatabase(config.GetConfig())
		})
	}
	return dbConn
}
