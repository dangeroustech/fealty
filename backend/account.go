package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//Account - Backend AccountID
type Account struct {
	AccountID    primitive.ObjectID `json:"accountid" bson:"_id"`
	RewardPoints int                `json:"rewardpoints"`
	Email        string             `json:"email"`
	Marketing    bool               `json:"marketing"`
}

//Accounts is also a thing
type Accounts []Account
