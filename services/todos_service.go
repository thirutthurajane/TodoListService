package services

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todoservice/configs"
	"todoservice/models"
)

func GetAllTodos() ([]models.Todo, error) {
	ctx, _ := configs.GetContext()
	client, err := configs.GetConn()
	if err != nil {
		return nil, errors.Wrap(err, "Configuration error")
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Connection error")
	}
	defer client.Disconnect(ctx)

	database := client.Database(configs.Configuration.Server.DbName)
	todoCollection := database.Collection("Todos")
	var todos []models.Todo
	cursor, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "Find collection error")
	}
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, errors.Wrap(err, "Cursor error")
	}
	return todos, nil
}

func GetTodo(id string) (todo []models.Todo, err error) {
	ctx, _ := configs.GetContext()
	client, err := configs.GetConn()

	if err != nil {
		return nil, errors.Wrap(err, "Configuration error")
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Connection error")
	}
	defer client.Disconnect(ctx)
	bsonId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "Matching object id error")
	}
	database := client.Database(configs.Configuration.Server.DbName)
	todosCollection := database.Collection("Todos")
	matchStage := bson.D{
		{"$match",
			bson.D{{"_id", bsonId}},
		},
	}
	cursor, err := todosCollection.Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		return nil, errors.Wrap(err, "Query Error")
	}
	if err = cursor.All(ctx, &todo); err != nil {
		return nil, errors.Wrap(err, "Cursor error")
	}
	return todo, nil
}
