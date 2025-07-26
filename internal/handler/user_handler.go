package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"go-rest/internal/model"
	"go-rest/internal/service"
)

func GetAllUsers(c *fiber.Ctx) error {
	log.Println("Getting all users")
	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	log.Println("Getting user", c.Params("userId"))

	// Convert userId string to uint
	userIdStr := c.Params("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	user, err := service.GetUser(uint(userId))
	if err != nil {
		// Check if it's a "record not found" error
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	log.Println("Creating user")
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	log.Println("Updating user", c.Params("userId"))

	// Convert userId string to uint
	userIdStr := c.Params("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	user, err := service.GetUser(uint(userId))
	if err != nil {
		// Check if it's a "record not found" error
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Parse the request body to update user data
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := service.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	log.Println("Deleting user", c.Params("userId"))

	// Convert userId string to uint
	userIdStr := c.Params("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	if err := service.DeleteUser(uint(userId)); err != nil {
		// Check if it's a "record not found" error
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}
