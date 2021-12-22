package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RenderAccounts - Allow Browser to View All Accounts
func adminAccounts(c *fiber.Ctx) error {
	a := MongoFindAll(50)

	// Render index template
	return c.Render("accounts_admin", fiber.Map{
		"Title":    "Accounts",
		"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
		"Accounts": a,
	})
}

// func searchAccounts(c *fiber.Ctx) error {
// 	a := MongoFindAll(50)
// 	// Render search template
// 	return c.Render("accounts_search", fiber.Map{"Accounts": a})
// }

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
	} else if result.Email == "EMPTY" {
		return c.JSON("{'Error': 'Account for This Email Already Exists'}")
	} else {
		return c.JSON(result)
	}
}

// CreateAccount - Create a New Account From The Admin Form
func createAccountForm(c *fiber.Ctx) error {
	var a Account
	if err := c.BodyParser(&a); err != nil {
		return c.Render("accounts_result", fiber.Map{
			"Message": fmt.Sprintf("Error: %s\nAccount Info: %#v", err, a),
			"Domain":  fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
		})
	}
	a.AccountID = primitive.NewObjectID()
	// return c.Render("accounts_result", fiber.Map{
	// 	"Message": fmt.Sprintf("Account Info: %#v", a),
	// 	"Domain":  fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
	// })
	result := MongoCreate(a)
	if result.Email == "DUPE" {
		return c.Render("accounts_result", fiber.Map{
			"Message": "Error: Account for This Email Already Exists",
			"Domain":  fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN"))})
	} else if result.Email == "EMPTY" {
		return c.Render("accounts_result", fiber.Map{
			"Message": fmt.Sprintf("Error: Email is Blank\n%#v", result),
			"Domain":  fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN"))})
	} else {
		return c.Render("accounts_result", fiber.Map{
			"Message": result,
			"Domain":  fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
		})
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
