package models

import (
	"Golang_Backend/db"
	"Golang_Backend/utils"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	User_id     string      `json:"user_id" validate:"required"`
}

func (todo *Todo) validate() error {
	return utils.Validate.Struct(todo)
}

func CreateTodo(todo *Todo) (*Todo, error) {

	if err := todo.validate(); err != nil {
		return nil, err
	}

	data, err := db.TodoModel.InsertOne(context.TODO(), todo)
	if err != nil {
		return nil, err
	}

	todo.ID = data.InsertedID

	return todo, nil

}

func UpdateTodo(todo *Todo, _id primitive.ObjectID, user_id primitive.ObjectID) error {

	if err := todo.validate(); err != nil {
		return fmt.Errorf("error while validating: %w", err)
	}

	_, err := db.TodoModel.UpdateOne(context.TODO(), bson.M{"_id": _id, "user_id": user_id.Hex()}, bson.M{"$set": todo})
	if err != nil {
		return fmt.Errorf("error while updating: %w", err)
	}

	data := db.TodoModel.FindOne(context.TODO(), bson.M{"_id": _id, "user_id": user_id.Hex()})
	if data.Err() == mongo.ErrNoDocuments {
		return fmt.Errorf("this document doesn't exist: %w", data.Err())
	} else if data.Err() != nil {
		return fmt.Errorf("error while fetching update: %w", data.Err())
	}

	if err := data.Decode(todo); err != nil {
		return fmt.Errorf("error while decoding: %w", err)
	}

	return nil

}
