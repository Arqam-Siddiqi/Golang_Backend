package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (u *User) Create(name string, email string, password string) (*User, error) {

	if err := u.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// user, err := db.UserModel.InsertOne(context.TODO(), bson.M{name: })
	return nil, nil

}
