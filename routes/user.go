package routes

import (
	"github.com/caiomp87/go-barber/controllers"
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(app *fiber.App) {
	app.Post("/user", controllers.CreateUser).Name("create-user")
	app.Get("/user", controllers.ListUsers).Name("list-users")
	app.Get("/user/:id", controllers.GetUser).Name("get-user")
	app.Patch("/user/:id", controllers.UpdateUser).Name("update-user")
	app.Delete("/user/:id", controllers.DeleteUser).Name("delete-user")
}
