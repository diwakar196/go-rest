package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-rest/internal/handler"
)

func SetupRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	userGroup := app.Group("/users")
	handler.RegisterRoutes(userGroup, db)

	return app
}
