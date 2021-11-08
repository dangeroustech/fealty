package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func renderAccounts(c *fiber.Ctx) error {
	a := MongoFindAll(50, "localhost")
	// Render index template
	return c.Render("accounts", fiber.Map{
		"Title":    "Accounts",
		"Accounts": a,
	})
}

func getAccounts(c *fiber.Ctx) error {
	a := MongoFindAll(50, "localhost")
	// Render index template
	return c.JSON(a)
}

func getAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)

	if err != nil {
		log.Println(err)
	}

	result := MongoFind(string(a.Email), "localhost")

	return c.JSON(result)
}

func createAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)
	a.AccountID = primitive.NewObjectID()

	if err != nil {
		log.Println(err)
	}

	result := MongoCreate(a, "localhost")

	return c.JSON(result)
}

func updateAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}

func deleteAccount(c *fiber.Ctx) error {

	return c.SendString(c.Params("accountId"))
}
