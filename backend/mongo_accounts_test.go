package main

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFind(t *testing.T) {
	result := MongoFind("000000", "localhost")

	if result.RewardPoints != 100 {
		t.Errorf("Account %s RewardPoints Incorrect. Expected %d, Got: %d.", "000000", 100, result.RewardPoints)
	}
}

func TestFindAll(t *testing.T) {
	result := MongoFindAll(200, "localhost")

	if result == nil {
		t.Errorf("Nil Data Returned for FindAll.")
	}
}

func TestCreate(t *testing.T) {
	id := primitive.NewObjectID()
	a := Account{id, 99, "test@test.com", true}
	result := MongoCreate(a, "localhost")

	if result.AccountID != id {
		t.Errorf("AccountID Created Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
	} else if result.RewardPoints != 99 {
		t.Errorf("RewardPoints Created Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
	} else if result.Email != "test@test.com" {
		t.Errorf("Email Created Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
	} else if result.Marketing != true {
		t.Errorf("Marketing Created Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
	}
}

func TestUpdate(t *testing.T) {
	id := primitive.NewObjectID()
	a := Account{id, 999, "test9@test.com", false}
	result := MongoUpdate(a, "localhost")

	if result.AccountID != id {
		t.Errorf("AccountID Updated Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
	} else if result.RewardPoints != 999 {
		t.Errorf("RewardPoints Updated Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
	} else if result.Email != "test9@test.com" {
		t.Errorf("Email Updated Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
	} else if result.Marketing != false {
		t.Errorf("Marketing Updated Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
	}
}
