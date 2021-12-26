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
	// Render admin interface
	return c.Render("accounts_admin", fiber.Map{
		"Title":    "Accounts",
		"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
		"Accounts": MongoFindAll(50),
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

	result := MongoFind(string(a.Email))
	if result.AccountID == primitive.NilObjectID {
		return c.JSON("{'Error': 'Account Not Found'}")
	} else {
		return c.JSON(result)
	}
}

type formEmail struct {
	Email string `form:"email"`
}

func getAccountForm(c *fiber.Ctx) error {
	var email formEmail
	if err := c.BodyParser(&email); err != nil {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %s\nEmail: %#v", err, email),
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
		})
	}

	result := MongoFind(string(email.Email))
	// check for not found
	if result.AccountID == primitive.NilObjectID {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Account For %s Not Found", email),
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
		})
	}

	return c.Render("accounts_result", fiber.Map{
		"Message":  result,
		"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
		"Accounts": MongoFindAll(50),
	})
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
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %s\nAccount Info: %#v", err, a),
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
		})
	}
	a.AccountID = primitive.NewObjectID()
	result := MongoCreate(a)
	if result.Email == "DUPE" {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  "Error: Account for This Email Already Exists",
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
		})
	} else if result.Email == "EMPTY" {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Email is Blank\n%#v", result),
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
		})
	} else {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  result,
			"Domain":   fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN")),
			"Accounts": MongoFindAll(50),
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

func updateAccountForm(c *fiber.Ctx) error {
	return nil
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

func deleteAccountForm(c *fiber.Ctx) error {
	return nil
}
