conn = new Mongo()
db = conn.getDB("fealty")

db.createCollection("accounts", {
    validator: { 
        $jsonSchema: {
            bsonType: "object",
            required: [ "accountid", "rewardpoints", "email", "marketing" ],
            properties: {
                accountid: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                rewardpoints: {
                    bsonType: "int",
                    description: "must be an integer and is required"
                },
                email: {
                    bsonType: "string",
                    description: "must be a string and is required"
                },
                marketing: {
                    bsonType: "bool",
                    description: "must be a bool and is required"
                },
            }
        }
    }
} )

db.createCollection("customers")
