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

type User struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string      `json:"name" validate:"required"`
	Email    string      `json:"email" validate:"required"`
	Password string      `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Validate() error {
	return utils.Validate.Struct(u)
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

func Login(email string, password string) (*User, error) {

	result := db.UserModel.FindOne(context.TODO(), bson.M{"email": email})
	if err := result.Err(); err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	} else if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("invalid credentials")
	}

	var user User
	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user result: %w", err)
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &user, nil

}
