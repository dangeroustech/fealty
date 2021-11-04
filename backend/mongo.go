package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoFind - Find An Account
func MongoFind(id string, server string) Account {
	// connect to mongo session running on localhost
	session, err := mgo.Dial(server)
	result := Account{}

	if err != nil {
		log.Print(err)
	}
	defer session.Close()

	// Set Mongo behaviour
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("wowrewards").C("accounts")

	// Execute the find
	err = c.Find(bson.M{"accountid": id}).One(&result)

	if err != nil {
		log.Print("Error finding account " + id)
		log.Print(err)
	} else {
		log.Printf("Account %s Searched For, Points %d Received.", result.AccountID, result.RewardPoints)
	}

	return result
}

// MongoFindAll - Return All Accounts
func MongoFindAll(limit int, server string) []Account {
	// connect to mongo session running on localhost
	session, err := mgo.Dial(server)
	result := []Account{}

	if err != nil {
		log.Print(err)
	}
	defer session.Close()

	// Set Mongo behaviour
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("wowrewards").C("accounts")

	// Execute The FindAll
	iter := c.Find(nil).Limit(limit).Iter()
	err = iter.All(&result)

	if err != nil {
		log.Print("Error getting accounts")
		log.Print(err)
	}

	return result
}

// MongoCreate - Create An Account
func MongoCreate(a Account, server string) Account {
	// connect to mongo session running on localhost
	session, err := mgo.Dial(server)

	if err != nil {
		log.Print(err)
	}
	defer session.Close()

	// Set Mongo behaviour
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("wowrewards").C("accounts")

	// Execute The Insert
	err = c.Insert(&Account{a.AccountID, a.RewardPoints, a.Email, a.Marketing})

	if err != nil {
		log.Printf("Error creating account %s", a.AccountID)
		log.Print(err)
	} else {
		log.Printf("Account: %s Created.", a.AccountID)
	}

	return a
}

// MongoUpdate - Update An Account
func MongoUpdate(a Account, server string) Account {
	// connect to mongo session running on localhost
	session, err := mgo.Dial(server)

	if err != nil {
		log.Print(err)
	}
	defer session.Close()

	// Set Mongo behaviour
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("wowrewards").C("accounts")

	// Execute The Update
	err = c.Update(bson.M{"accountid": a.AccountID}, a)

	if err != nil {
		log.Print("Error updating account " + a.AccountID)
		log.Print(err)
	} else {
		log.Printf("Account %s Updated. Reward Points: %d, Email: %s, Marketing: %t",
			a.AccountID, a.RewardPoints, a.Email, a.Marketing)
	}

	return a
}
