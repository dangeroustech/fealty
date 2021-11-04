package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	return c.Next()
}

func main() {
	// Set Up Fiber App
	app := fiber.New()
	app.Use(logger.New())

	// Root API Route
	api := app.Group("/api", middleware) // /api

	// API v1 Routes
	v1 := api.Group("/v1", middleware) // /api/v1
	v1.Get("/accounts", getAccounts)
	v1.Get("/account/{accountId}", getAccount)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("You took a wrong turn there, adventurer!")
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
