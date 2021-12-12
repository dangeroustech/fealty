package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func middleware(c *fiber.Ctx) error {
	fmt.Printf("Request from %#v\n", c.IP()) // technically only needs %s
	return c.Next()
}

func main() {
	// Initialize standard Go html template engine
	htmlEngine := html.New(os.Getenv("FEALTY_CONFIG")+"/static", ".html")

	// Set Up Fiber App
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "FealTY API v1",
		AppName:       "FealTY v0.5.0",
		Views:         htmlEngine,
	})
	// Logging
	app.Use(logger.New())
	// Auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendStatus(401)
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	// Root API Route
	api := app.Group("/api") // /api

	// API v1 Routes
	v1 := api.Group("/v1") // /api/v1

	// Mass Account Routes
	v1.Get("/accounts/get", getAccounts, middleware)     // /api/v1/accounts
	v1.Get("/accounts/view", renderAccounts, middleware) // /api/v1/accounts/view

	// Individual Account Routes
	acc := v1.Group("/account", middleware) // /api/v1/account
	acc.Get("", getAccount)
	acc.Post("", createAccount)
	acc.Put("", updateAccount)
	acc.Delete("", deleteAccount)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("You took a wrong turn there, adventurer!")
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
