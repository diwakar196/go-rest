package main

import (
	"log"

	"go-rest/internal/api"
	"go-rest/internal/model"
	"go-rest/pkg/config"
	"go-rest/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connect(cfg)
	db.AutoMigrate(&model.User{})

	app := api.SetupRouter(db)
	log.Printf("Server is running on port : %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
