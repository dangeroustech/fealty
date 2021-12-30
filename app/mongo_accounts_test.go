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

func TestMain(m *testing.M) {
	MongoCreate(a)
	m.Run()
	MongoDelete(a)
}

func TestFind(t *testing.T) {
	result := MongoFind("test@test.com")

	if result.RewardPoints != 100 {
		t.Errorf("Account for %s RewardPoints Incorrect. Expected %d, Got: %d.", "test@test.com", 100, result.RewardPoints)
	}
}

func TestFindAll(t *testing.T) {
	result := MongoFindAll(1)

	if result == nil {
		t.Errorf("Nil Data Returned for FindAll.")
	}
}

func TestUpdate(t *testing.T) {
	// account details to update
	newA := a
	newA.RewardPoints = 200

	// update a with newA details
	result := MongoUpdate(newA)

	if result.AccountID != newA.AccountID {
		t.Errorf("AccountID Updated Incorrectly. Expected %s, Got %s.", newA.AccountID, result.AccountID)
	} else if result.RewardPoints != newA.RewardPoints {
		t.Errorf("RewardPoints Updated Incorrectly. Expected %d, Got %d.", newA.RewardPoints, result.RewardPoints)
	} else if result.Email != newA.Email {
		t.Errorf("Email Updated Incorrectly. Expected %s, Got %s.", newA.Email, result.Email)
	} else if result.Marketing != newA.Marketing {
		t.Errorf("Marketing Updated Incorrectly. Expected %t, Got %t.", newA.Marketing, result.Marketing)
	}
}
