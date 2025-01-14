package controllers

import (
	"Golang_Backend/db"
	"Golang_Backend/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAll(ctx *gin.Context) {

	cursor, err := db.UserModel.Find(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	defer cursor.Close(context.TODO())

	var users []bson.M
	if err := cursor.All(context.TODO(), &users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, users)
}

func DeleteUserByJwt(ctx *gin.Context) {

	_id, _ := ctx.Get("user_id")

	result := db.UserModel.FindOneAndDelete(context.TODO(), bson.M{"_id": _id})

	if result.Err() != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "This Object doesn't exist."})
		ctx.Abort()
		return
	}

	if _, err := db.TodoModel.DeleteMany(context.TODO(), bson.M{"user_id": _id.(primitive.ObjectID).Hex()}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting Associated Todos."})
		ctx.Abort()
		return
	}

	var user bson.M
	if err := result.Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding deleted user."})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func FindById(ctx *gin.Context) {

	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Object ID."})
		return
	}

	result := db.UserModel.FindOne(context.TODO(), bson.M{"_id": _id})
	if result.Err() != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Err().Error()})
		return
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func GetUserByJwt(ctx *gin.Context) {

	_id, _ := ctx.Get("user_id")
	var user models.User
	if err := db.UserModel.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func UpdateUserById(ctx *gin.Context) {

	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_user, err := models.UpdateUser(&user, _id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, new_user)

}
