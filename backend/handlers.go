package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

//Index does index things
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "These are not the droids you're looking for...")
}

func getAccounts(c *fiber.Ctx) error {
	return c.SendString(c.Path())
}

func getAccount(c *fiber.Ctx) error {
	return c.SendString(c.Path())
}
