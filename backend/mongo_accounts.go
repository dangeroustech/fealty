package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func dbConnect() *mongo.Client {
	MONGO_URI := os.Getenv("MONGODB_URI")
	MONGO_USER := os.Getenv("MONGODB_USER")
	MONGO_PASS := os.Getenv("MONGODB_PASS")
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(MONGO_URI).SetAuth(options.Credential{Username: MONGO_USER, Password: MONGO_PASS})
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// MongoFind - Find An Account
func MongoFind(email string, server string) Account {
	// connect to mongo session running on localhost
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")
	var a Account

	// Execute the find
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&a)

	if err != nil {
		log.Printf("Error finding account for %s", email)
		log.Print(err)
	} else {
		log.Printf("Account %s For %s Has %d Points.", a.AccountID, a.Email, a.RewardPoints)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	return a
}

// MongoFindAll - Return All Accounts
func MongoFindAll(limit int64, server string) []*Account {
	// connect to mongo session running on localhost
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")
	var results []*Account

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	// Execute The FindAll
	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Account
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	return results
}

// MongoCreate - Create An Account
func MongoCreate(a Account, server string) Account {
	// connect to mongo session running on localhost
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")

	// Execute The Insert
	result, err := collection.InsertOne(context.TODO(), a)

	if err != nil {
		log.Printf("Error creating account %s", a.AccountID)
		log.Print(err)
	} else {
		log.Printf("Account %s Created at %s", a.AccountID, result.InsertedID)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
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
	c := session.DB("fealty").C("accounts")

	// Execute The Update
	err = c.Update(bson.M{"accountid": a.AccountID}, a)

	if err != nil {
		log.Printf("Error updating account %v", a.AccountID)
		log.Print(err)
	} else {
		log.Printf("Account %s Updated. Reward Points: %d, Email: %s, Marketing: %t",
			a.AccountID, a.RewardPoints, a.Email, a.Marketing)
	}

	return a
}
