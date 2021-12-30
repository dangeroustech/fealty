package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Determine correct domain
func GetDomain() string {
	if os.Getenv("FEALTY_ENV") == "TEST" {
		return os.Getenv("DOMAIN")
	}
	return fmt.Sprintf("rewards.%s", os.Getenv("DOMAIN"))
}

func healthcheck(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// adminAccounts - Allow Browser to View All Admin Interface
func adminAccounts(c *fiber.Ctx) error {
	// Render admin interface
	return c.Render("accounts_admin", fiber.Map{
		"Message":  "",
		"Domain":   GetDomain(),
		"Accounts": MongoFindAll(50),
		"Error":    0,
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

func getAccountForm(c *fiber.Ctx) error {
	var email FormEmail
	if err := c.BodyParser(&email); err != nil {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %s\nEmail: %s", err, email.Email),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}

	result := MongoFind(email.Email)
	// check for not found
	if result.AccountID == primitive.NilObjectID {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Account For '%s' Not Found", email.Email),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}

	return c.Render("accounts_admin", fiber.Map{
		"Message":  result,
		"Domain":   GetDomain(),
		"Accounts": MongoFindAll(50),
		"Error":    0,
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
	// var a Account
	a := new(Account)
	if err := c.BodyParser(a); err != nil {
		errVal := ValidateAccount(*a)
		if errVal != nil {
			return c.Render("accounts_admin", fiber.Map{
				"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, a),
				"Domain":   GetDomain(),
				"Accounts": MongoFindAll(50),
				"Error":    1,
			})
		}
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", err, a),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	a.AccountID = primitive.NewObjectID()
	errVal := ValidateAccount(*a)
	if errVal != nil {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, a),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	result := MongoCreate(*a)
	if result.Email == "DUPE" {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  "Error: Account for This Email Already Exists",
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	} else if result.Email == "EMPTY" {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Email is Blank\n%#v", result),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	} else {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  result,
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    0,
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
		return c.JSON("{'Error': 'Account Not Found'}")
	} else {
		return c.JSON(result)
	}
}

// UpdateAccount - Update a New Account From The Admin Form
func updateAccountForm(c *fiber.Ctx) error {
	a := new(Account)
	if err := c.BodyParser(a); err != nil {
		errVal := ValidateAccount(*a)
		if errVal != nil {
			return c.Render("accounts_admin", fiber.Map{
				"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, a),
				"Domain":   GetDomain(),
				"Accounts": MongoFindAll(50),
				"Error":    1,
			})
		}
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", err, a),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	errVal := ValidateAccount(*a)
	if errVal != nil {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, a),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	a.AccountID = MongoFind(a.Email).AccountID
	result := MongoUpdate(*a)
	if result.AccountID == primitive.NilObjectID {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Account Not Found\nSearch: %#v - Result: %#v", a, result),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	} else {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  result,
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    0,
		})
	}
}

// DeleteAccount - Delete an Account
func deleteAccount(c *fiber.Ctx) error {
	var a Account
	err := json.Unmarshal(c.Body(), &a)

	if err != nil {
		log.Println(err)
	}

	result := MongoDelete(MongoFind(a.Email))

	if result.AccountID == primitive.NilObjectID {
		return c.JSON("{'Error': 'Account Not Found'")
	} else {
		return c.JSON(result)
	}
}

func deleteAccountForm(c *fiber.Ctx) error {
	email := new(FormEmail)
	if err := c.BodyParser(email); err != nil {
		errVal := ValidateEmail(*email)
		if errVal != nil {
			return c.Render("accounts_admin", fiber.Map{
				"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, email),
				"Domain":   GetDomain(),
				"Accounts": MongoFindAll(50),
				"Error":    1,
			})
		}
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", err, email),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	errVal := ValidateEmail(*email)
	if errVal != nil {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: %#v\nAccount Info: %#v", errVal, email),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	}
	result := MongoDelete(MongoFind(email.Email))
	if result.AccountID == primitive.NilObjectID {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  fmt.Sprintf("Error: Account Not Found\nSearch: %#v - Result: %#v", email, result),
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    1,
		})
	} else {
		return c.Render("accounts_admin", fiber.Map{
			"Message":  result,
			"Domain":   GetDomain(),
			"Accounts": MongoFindAll(50),
			"Error":    0,
		})
	}
}
