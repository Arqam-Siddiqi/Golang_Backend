package models

import (
	"Golang_Backend/db"
	"Golang_Backend/utils"
	"context"
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
