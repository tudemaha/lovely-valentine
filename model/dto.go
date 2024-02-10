package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}

type NewMessage struct {
	Magic       string   `bson:"magic"`
	HashedMagic string   `bson:"hashed_magic"`
	Quotes      []string `bson:"quotes"`
}

type GetMessage struct {
	ID primitive.ObjectID `bson:"_id"`
	NewMessage
}
