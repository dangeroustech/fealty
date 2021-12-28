package main

import "go.mongodb.org/mongo-driver/bson/primitive"

//Account - Backend AccountID
type Account struct {
	AccountID    primitive.ObjectID `json:"accountid" bson:"_id"`
	RewardPoints int                `json:"rewardpoints" form:"points" validate:"required,number"`
	Email        string             `json:"email" form:"email" validate:"required,email"`
	Marketing    bool               `json:"marketing" form:"marketing"`
}

//Accounts is also a thing
type Accounts []Account
