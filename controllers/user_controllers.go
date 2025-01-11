package controllers

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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if _, err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUserById(c *gin.Context) {

	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Object ID"})
		return
	}

	result := db.UserModel.FindOneAndDelete(context.TODO(), bson.M{"_id": _id})

	if result.Err() != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "This Object doesn't exist."})
		return
	}

	var user bson.M
	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding deleted user:"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func FindById(c *gin.Context) {

	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Object ID"})
		return
	}

	result := db.UserModel.FindOne(context.TODO(), bson.M{"_id": _id})
	if result.Err() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Err().Error()})
		return
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserById(c *gin.Context) {

	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_user, err := models.UpdateUser(&user, _id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, new_user)

}
