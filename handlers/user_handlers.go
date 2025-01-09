package handlers

import (
	"Golang_Backend/db"
	"Golang_Backend/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll(c *gin.Context) {
	cursor, err := db.UserModel.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var users []bson.M
	if err := cursor.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if _, err := db.UserModel.InsertOne(context.TODO(), user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserById(c *gin.Context) {

	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Object ID"})
		return
	}

	result := db.UserModel.FindOneAndDelete(context.TODO(), bson.M{"_id": _id})

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "This Object doesn't exist."})
		return
	}

	var user bson.M
	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error decoding deleted user:"})
		return
	}

	c.JSON(http.StatusOK, user)

}
