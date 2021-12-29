package main

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var a = Account{
	AccountID:    primitive.NewObjectIDFromTimestamp(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),
	Email:        "test@test.com",
	RewardPoints: 100,
	Marketing:    false,
}

func TestFind(t *testing.T) {
	TestPrep(a)
	result := MongoFind("test@test.com")

	if result.RewardPoints != 100 {
		t.Errorf("Account for %s RewardPoints Incorrect. Expected %d, Got: %d.", "test@test.com", 100, result.RewardPoints)
	}

	TestCleanup(a)
}

func TestFindAll(t *testing.T) {
	TestPrep(a)
	result := MongoFindAll(1)

	if result == nil {
		t.Errorf("Nil Data Returned for FindAll.")
	}

	TestCleanup(a)
}

func TestCreate(t *testing.T) {
	result := MongoCreate(a)

	if result.AccountID != a.AccountID {
		t.Errorf("AccountID Created Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
	} else if result.RewardPoints != a.RewardPoints {
		t.Errorf("RewardPoints Created Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
	} else if result.Email != a.Email {
		t.Errorf("Email Created Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
	} else if result.Marketing != a.Marketing {
		t.Errorf("Marketing Created Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
	}

	TestCleanup(a)
}

func TestUpdate(t *testing.T) {
	// create Account a
	TestPrep(a)

	// create a new Account with the same ID as a
	newA := Account{AccountID: a.AccountID, Email: "test1@test1.com", RewardPoints: 350, Marketing: true}

	// update a with newA details
	result := MongoUpdate(newA)

	if result.AccountID != newA.AccountID {
		t.Errorf("AccountID Updated Incorrectly. Expected %s, Got %s.", a.AccountID, result.AccountID)
	} else if result.RewardPoints != newA.RewardPoints {
		t.Errorf("RewardPoints Updated Incorrectly. Expected %d, Got %d.", a.RewardPoints, result.RewardPoints)
	} else if result.Email != newA.Email {
		t.Errorf("Email Updated Incorrectly. Expected %s, Got %s.", a.Email, result.Email)
	} else if result.Marketing != newA.Marketing {
		t.Errorf("Marketing Updated Incorrectly. Expected %t, Got %t.", a.Marketing, result.Marketing)
	}
	TestCleanup(newA)
}

func TestDelete(t *testing.T) {
	TestPrep(a)

	result := MongoDelete(a)

	if result.AccountID != a.AccountID {
		t.Errorf("Account Deleted Incorrectly. Expected %s, Got %s", a.AccountID, result.AccountID)
	}
}
