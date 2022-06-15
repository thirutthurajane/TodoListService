package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Detail     string             `bson:"detail"`
	IsComplete bool               `bson:"isComplete"`
	Created    time.Time          `bson:"created"`
	User       primitive.ObjectID `bson:"user"`
}
