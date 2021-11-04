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
	a := Accounts{
		Account{
			AccountID:    "1234",
			RewardPoints: 10,
			Email:        "test@test.com",
			Marketing:    false,
		},
		Account{
			AccountID:    "5678",
			RewardPoints: 100,
			Email:        "test1@test.com",
			Marketing:    true,
		},
	}
	// Render index template
	return c.Render("default", fiber.Map{
		"Title":    "Accounts",
		"Accounts": a,
	})
	// return c.JSON(a)
}

func getAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}

func createAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}

func updateAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}

func deleteAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}
