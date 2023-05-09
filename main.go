package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caiomp87/go-barber/middlewares"
	"github.com/caiomp87/go-barber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ReadBufferSize:        8 * 2068,
		StrictRouting:         true,
		CaseSensitive:         true,
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
	})

	routes.AddHealthRoute(app)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(middlewares.Authenticate)

	routes.AddUserRoutes(app)

	go func() {
		log.Println("API listening on port 3000")
		app.Listen(":3000")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		log.Fatal("failed to shutdown gracefully")
	}

	log.Println("API shutdown gracefully")
}
