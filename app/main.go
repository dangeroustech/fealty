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

// func middleware(c *fiber.Ctx) error {
// 	fmt.Printf("Request from %#v\n", c.IP()) // technically only needs %s
// 	auth := fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("FEALTY_USER"), os.Getenv("FEALTY_PASS")))))
// 	c.Set("Authorization", auth)
// 	return c.Next()
// }

func AuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			os.Getenv("FEALTY_USER"): os.Getenv("FEALTY_PASS"),
		},
		Realm: "Admin",
		Unauthorized: func(c *fiber.Ctx) error {
			c.Set("WWW-Authenticate", "Basic realm=\"Admin\"")
			return c.SendStatus(401)
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}
	err := basicauth.New(cfg)
	return err
}

func main() {
	// Initialize standard Go html template engine
	htmlEngine := html.New(os.Getenv("FEALTY_CONFIG")+"/static", ".html")
	htmlEngine.AddFunc("updateAccount", updateAccount)
	htmlEngine.AddFunc("deleteAccount", deleteAccount)

	// Set Up Fiber App
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "FealTY API v1",
		AppName:       "FealTY v0.6.0",
		Views:         htmlEngine,
	})

	// Logging
	app.Use(logger.New())

	// Static Files
	app.Static("/static", os.Getenv("FEALTY_CONFIG")+"/static")

	// Root API Route
	api := app.Group("/api") // /api

	// API v1 Routes
	v1 := api.Group("/v1") // /api/v1

	// Mass Account Routes
	accs := v1.Group("/accounts")                // /api/v1/accounts
	accs.Get("/admin", AuthReq(), adminAccounts) // /api/v1/accounts/view
	accs.Get("/get", AuthReq(), getAccounts)     // /api/v1/accounts/get

	// Individual Account Routes
	acc := v1.Group("/account", AuthReq()) // /api/v1/account
	acc.Get("", getAccount)
	acc.Post("", createAccount)
	acc.Put("", updateAccount)
	acc.Delete("", deleteAccount)
	acc.Post("/form", createAccountForm)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("You took a wrong turn there, adventurer!")
		return c.SendStatus(404)
		// => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
