package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-rest/internal/model"
	"go-rest/internal/repository"
	"go-rest/internal/service"
)

func RegisterRoutes(r fiber.Router, db *gorm.DB) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	r.Post("/", func(c *fiber.Ctx) error { return CreateUser(c, service) })
	r.Get("/", func(c *fiber.Ctx) error { return GetAllUsers(c, service) })
}

func GetAllUsers(c *fiber.Ctx, service *service.Service) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx, service *service.Service) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
