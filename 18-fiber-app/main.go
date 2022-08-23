package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	controller "github.com/jabhishek87/fiberapp/handler"
	"github.com/jabhishek87/fiberapp/model"
)

const port = "4000"

func main() {
	fmt.Println("Golang Fiber app")

	// Seed records
	courses := model.CreateSeedRecords()
	fmt.Println(courses)

	// init fiber App
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("Main route")
		message := "Welcome to Fiber App on Port: " + port
		return ctx.Status(fiber.StatusOK).SendString(message)
	})

	// books.RegisterRoutes(app, db)
	controller.RegisterRoutes(app, courses)

	app.Listen(":" + port)

}
