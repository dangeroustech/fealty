package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func AuthSource() string {
	if os.Getenv("FEALTY_ENV") == "TEST" {
		return "admin"
	}
	return "fealty"
}

func dbConnect() *mongo.Client {
	MONGO_URI := os.Getenv("MONGO_URI")
	MONGO_USER := os.Getenv("MONGO_USER")
	MONGO_PASS := os.Getenv("MONGO_PASS")
	AUTH_SOURCE := AuthSource()
	// log.Printf("Connecting to Mongo at %s (authSource %s) with username %s and password %s", MONGO_URI, AUTH_SOURCE, MONGO_USER, MONGO_PASS)
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(MONGO_URI).SetAuth(options.Credential{Username: MONGO_USER, Password: MONGO_PASS, AuthSource: AUTH_SOURCE})
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
func MongoFind(email string) Account {
	// connect to mongo session
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")
	var a Account

	// Execute the find
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&a)

	if err != nil {
		log.Printf("Error finding account for %s:\n%#v", email, err)
		a.AccountID = primitive.NilObjectID
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found account %v for provided email %s", a.AccountID, email)

	return a
}

// MongoFindAll - Return All Accounts
func MongoFindAll(limit int64) []*Account {
	// connect to mongo session
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

	return results
}

// MongoCreate - Create An Account
func MongoCreate(a Account) Account {
	// connect to mongo session
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")

	// Validate
	if a.Email == "" {
		log.Printf("Error creating account %s", a.AccountID)
		a.Email = "EMPTY"
		return a
	}

	// Check for Duplicate
	if MongoFind(a.Email).Email != "" {
		a.Email = "DUPE"
		return a
	}

	// Execute The Insert
	result, err := collection.InsertOne(context.TODO(), a)

	if err != nil {
		log.Printf("Error creating account %s", a.AccountID)
		log.Print(err)
	} else {
		log.Printf("1 Account(s) Created (%s). Reward Points: %d, Email: %s, Marketing: %t",
			result.InsertedID, a.RewardPoints, a.Email, a.Marketing)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	return a
}

// MongoUpdate - Update An Account
func MongoUpdate(a Account) Account {
	// connect to mongo session
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")
	log.Printf("a @ MongoUpdate: %#v", a)

	// Execute The Update
	result, _ := collection.ReplaceOne(context.TODO(), bson.M{"email": a.Email}, a)

	if result.MatchedCount == 0 {
		log.Printf("Error updating account %v with payload %#v", a.AccountID, a)
	} else {
		log.Printf("%v Account(s) Updated (%s). Reward Points: %d, Email: %s, Marketing: %t",
			result.ModifiedCount, a.AccountID, a.RewardPoints, a.Email, a.Marketing)
		log.Printf("RESULT: %#v", result)
	}

	if result.ModifiedCount == 1 {
		return a
	} else {
		a.AccountID = primitive.NilObjectID
		return a
	}
}

// MongoDelete - Delete An Account
func MongoDelete(a Account) Account {
	// connect to mongo session
	client := dbConnect()
	collection := client.Database("fealty").Collection("accounts")

	// Execute The Deletion
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": a.AccountID})
	// log.Printf("RESULT: %#v", result)

	if err != nil {
		log.Printf("Error while deleting: %v", err)
	}

	// if for some reason we didn't delete anything
	// even though the account existed
	if result.DeletedCount == 0 {
		a.AccountID = primitive.NilObjectID
		return a
	} else {
		log.Printf("%d Account(s) Deleted (%s). Reward Points: %d, Email: %s, Marketing: %t",
			result.DeletedCount, a.AccountID, a.RewardPoints, a.Email, a.Marketing)
		return a
	}
}
