package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func middleware(c *fiber.Ctx) error {
	fmt.Printf("Request from %#v\n", c.IP()) // technically only needs %s
	return c.Next()
}

func main() {
	// Initialize standard Go html template engine
	htmlEngine := html.New("./static", ".html")

	// Set Up Fiber App
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "FealTY API v1",
		AppName:       "FealTY v0.0.1",
		Views:         htmlEngine,
	})
	app.Use(logger.New())

	// Root API Route
	api := app.Group("/api", middleware) // /api

	// API v1 Routes
	v1 := api.Group("/v1", middleware) // /api/v1

	// Account Routes
	v1.Get("/accounts", getAccounts)

	acc := v1.Group("/account", middleware)
	acc.Get("/:accountId", getAccount)
	acc.Post("/:accountId", createAccount)
	acc.Put("/:accountId", updateAccount)
	acc.Delete("/:accountId", deleteAccount)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("You took a wrong turn there, adventurer!")
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
