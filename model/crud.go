package model

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(message NewMessage) (string, error) {
	conn := MongoConnection(os.Getenv("MONGO_COLLECTION"))
	defer MongoCloseConnection(conn)

	result, err := conn.Coll.InsertOne(context.TODO(), message)
	if err != nil {
		return "", err
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedId, nil
}

func GetMessage(id string) (Message, error) {
	var message Message

	conn := MongoConnection(os.Getenv("MONGO_COLLECTION"))
	defer MongoCloseConnection(conn)

	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	if err := conn.Coll.FindOne(context.TODO(), filter).Decode(&message); err != nil {
		return message, err
	}

	return message, nil
}
