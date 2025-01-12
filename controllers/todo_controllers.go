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

func GetAllTodos(ctx *gin.Context) {

	cursor, err := db.TodoModel.Find(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	defer cursor.Close(context.TODO())

	var todos []models.Todo
	if err := cursor.All(context.TODO(), &todos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, todos)

}

func CreateTodo(ctx *gin.Context) {

	_id, _ := ctx.Get("user_id")

	var todo models.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	todo.User_id = _id.(primitive.ObjectID).Hex()

	if _, err := models.CreateTodo(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, todo)

}

func GetUserTodos(ctx *gin.Context) {

	_id, _ := ctx.Get("user_id")
	result, err := db.TodoModel.Find(context.TODO(), bson.M{"user_id": _id.(primitive.ObjectID).Hex()})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	defer result.Close(context.TODO())

	var todos []models.Todo
	if err := result.All(context.TODO(), &todos); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, todos)
}
