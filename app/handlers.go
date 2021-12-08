package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RenderAccounts - Allow Browser to View All Accounts
func renderAccounts(c *fiber.Ctx) error {
	a := MongoFindAll(50)
	// Render index template
	return c.Render("accounts", fiber.Map{
		"Title":    "Accounts",
		"Domain":   "THISISADOMAIN",
		"Accounts": a,
	})
}

// GetAccounts - API Query to Return All Accounts as JSON
func getAccounts(c *fiber.Ctx) error {
	// Get All Accounts
	a := MongoFindAll(50)

	// Render index template
	return c.JSON(a)
}

// GetAccount - Get a Single Account With Email
func getAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)

	if err != nil {
		log.Println(err)
	}

	result := MongoFind(string(a.Email), false)
	if result.AccountID == primitive.NilObjectID {
		return c.JSON("{'Error': 'Account Not Found'}")
	} else {
		return c.JSON(result)
	}
}

// CreateAccount - Create a New Account
func createAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)
	a.AccountID = primitive.NewObjectID()

	if err != nil {
		log.Println(err)
	}

	result := MongoCreate(a)
	if result.Email == "DUPE" {
		return c.JSON("{'Error': 'Account for This Email Already Exists'}")
	} else {
		return c.JSON(result)
	}
}

// UpdateAccount - Update an Account With Email
func updateAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)

	if err != nil {
		log.Println(err)
	}

	result := MongoUpdate(a)
	if result.AccountID == primitive.NilObjectID {
		return c.JSON("{'Error': 'Account Not Found'")
	} else {
		return c.JSON(result)
	}
}

// DeleteAccount - Delete an Account
func deleteAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)

	if err != nil {
		log.Println(err)
	}

	log.Printf("%#v", a)
	result := MongoDelete(a.Email)

	if result.AccountID == primitive.NilObjectID {
		return c.JSON("{'Error': 'Account Not Found'")
	} else {
		return c.JSON(result)
	}
}
