package models

import (
	"Golang_Backend/db"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string      `json:"name" validate:"required"`
	Email    string      `json:"email" validate:"required"`
	Password string      `json:"password" validate:"required"`
}

var validate = validator.New()

func (u *User) Validate() error {
	return validate.Struct(u)
}

func CreateUser(u *User) (*User, error) {

	if err := u.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	doc, err := bson.Marshal(u)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user to BSON: %w", err)
	}

	user, err := db.UserModel.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user into collection: %w", err)
	}

	u.ID = user.InsertedID

	return u, nil

}

func UpdateUser(u *User, _id primitive.ObjectID) (*User, error) {

	if err := u.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	update := bson.M{
		"$set": u, // This will update only the fields in the struct
	}

	_, err := db.UserModel.UpdateByID(context.TODO(), _id, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return u, nil

}
