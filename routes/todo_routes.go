package routes

import (
	"Golang_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.RouterGroup) {

	router.GET("/", controllers.GetAllTodos)
	router.POST("/", controllers.CreateTodo)
	router.GET("/self", controllers.GetUserTodos)

}
