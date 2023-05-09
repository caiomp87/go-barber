package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddHealthRoute(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		fmt.Println(c.Method() + c.App().GetRoute(c.Route().Name).Path)

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "service is running successfully",
		})
	}).Name("health")
}
