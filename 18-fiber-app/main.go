package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jabhishek87/fiberapp/controller"
	"github.com/jabhishek87/fiberapp/model"
)

const port = "4000"

func main() {

	fmt.Println("Golang Fiber app")
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// start the application on http://localhost:3000
	log.Fatal(app.Listen(":" + port))
}

func Setup() *fiber.App {

	// Seed records
	courses := model.CreateSeedRecords()
	fmt.Println(courses)

	// Initialize a new app
	app := fiber.New()

	// Register the index route with a simple
	// "OK" response. It should return status
	// code 200
	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Main route")
		message := "Welcome to Fiber App on Port: " + port
		return ctx.Status(fiber.StatusOK).SendString(message)
	})
	controller.RegisterRoutes(app, courses)
	// Return the configured app
	return app
}
