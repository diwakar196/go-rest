package main

import (
	"log"

	"github.com/diwakar196/go-rest/api"
	"github.com/diwakar196/go-rest/pkg/config"
	"github.com/diwakar196/go-rest/pkg/database"
)

func main() {
	cfg := config.LoadConfig()
	db := database.Connet(cfg)

	r := api.SetupRouter(db)
	log.Println("Server is running on port : %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
