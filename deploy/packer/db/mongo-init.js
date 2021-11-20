conn = new Mongo()

db.createUser({
    user: process.env.MONGODB_ROOT_USER,
    pwd:  process.env.MONGODB_ROOT_PASS,
    roles: [{role: "root" , db: "admin"}]
})

db = conn.getDB("fealty")

db.createUser({
    user: "fealty",
    pwd:  process.env.MONGODB_FEALTY_PASS,
    roles: [{role: "dbOwner" , db: "fealty"}]
})

db.createCollection("accounts", {
    validator: { 
        $jsonSchema: {
            bsonType: "object",
            required: [ "rewardpoints", "email", "marketing" ],
            properties: {
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
