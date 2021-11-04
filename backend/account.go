package main

//Account - Backend AccountID
type Account struct {
	AccountID    string `json:"accountid"`
	RewardPoints int    `json:"rewardpoints"`
	Email        string `json:"email"`
	Marketing    bool   `json:"marketing"`
}

//Accounts is also a thing
type Accounts []Account
