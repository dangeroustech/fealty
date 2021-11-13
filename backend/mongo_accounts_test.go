package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var a = Account{AccountID: primitive.NilObjectID, Email: "test@test.com", RewardPoints: 100, Marketing: true}

func TestFind(t *testing.T) {
	TestPrep(a)
	t.Log("Prepped For TestFind")
	result := MongoFind("test@test.com", "localhost")

	if result.RewardPoints != 100 {
		t.Errorf("Account for %s RewardPoints Incorrect. Expected %d, Got: %d.", "test@test.com", 100, result.RewardPoints)
	}

	TestCleanup(a.Email)
	t.Log("Cleaned Up After TestFind")
}

func TestFindAll(t *testing.T) {
	TestPrep(a)
	t.Log("Prepped For TestFindAll")
	result := MongoFindAll(1, "localhost")

	if result == nil {
		t.Errorf("Nil Data Returned for FindAll.")
	}

	TestCleanup(a.Email)
	t.Log("Cleaned Up After TestFindAll")
}

// func TestCreate(t *testing.T) {
// 	id := primitive.NewObjectID()
// 	a := Account{id, 99, "test@test.com", true}
// 	result := MongoCreate(a, "localhost")

// 	if result.AccountID != id {
// 		t.Errorf("AccountID Created Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
// 	} else if result.RewardPoints != 99 {
// 		t.Errorf("RewardPoints Created Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
// 	} else if result.Email != "test@test.com" {
// 		t.Errorf("Email Created Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
// 	} else if result.Marketing != true {
// 		t.Errorf("Marketing Created Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	id := primitive.NewObjectID()
// 	a := Account{id, 999, "test9@test.com", false}
// 	result := MongoUpdate(a, "localhost")

// 	if result.AccountID != id {
// 		t.Errorf("AccountID Updated Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
// 	} else if result.RewardPoints != 999 {
// 		t.Errorf("RewardPoints Updated Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
// 	} else if result.Email != "test9@test.com" {
// 		t.Errorf("Email Updated Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
// 	} else if result.Marketing != false {
// 		t.Errorf("Marketing Updated Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
// 	}
// }

// func TestDelete(t *testing.T) {
// }
