package model

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(message NewMessage) (string, bool) {

	conn := MongoConnection(os.Getenv("MONGO_COLLECTION"))
	defer MongoCloseConnection(conn)

	result, err := conn.Coll.InsertOne(context.TODO(), message)
	if err != nil {
		log.Printf("ERROR CreateMessage fatal error: %v", err)
		return "", false
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedId, true
}
