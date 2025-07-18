package main

import (
	"log"

	"go-rest/internal/api"
	"go-rest/internal/model"
	"go-rest/pkg/config"
	"go-rest/pkg/database"

	"gorm.io/gorm"
)

func main() {
	config := config.GetConfig()
	db := database.GetDB()
	migrateDatabase(db)

	app := api.FinanceManagementApp(db)
	log.Printf("Server is running on port : %s", config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
