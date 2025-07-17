package main

import (
	"log"

	"go-rest/api"
	"go-rest/model"
	"go-rest/pkg/config"
	"go-rest/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	db.AutoMigrate(&model.User{})

	r := api.SetupRouter(db)
	log.Printf("Server is running on port : %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
