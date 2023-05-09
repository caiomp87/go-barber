package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
}

func ListUsers(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "users listed successfully",
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user found successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user updated successfully",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
