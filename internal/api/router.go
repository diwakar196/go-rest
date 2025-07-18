package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-rest/internal/handler"
)

func FinanceManagementApp(db *gorm.DB) *fiber.App {
	financeMgmtApp := fiber.New()
	apiRoutes := financeMgmtApp.Group("/api")
	registerPing(apiRoutes)
	registerUserRoutes(apiRoutes)
	return financeMgmtApp
}

func registerPing(apiRoutes fiber.Router) {
	apiRoutes.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy!")
	})
}

func registerUserRoutes(apiRoutes fiber.Router) {
	userRoutes := apiRoutes.Group("/users")
	userRoutes.Get("/", handler.GetAllUsers)
	userRoutes.Post("/", handler.CreateUser)
}
