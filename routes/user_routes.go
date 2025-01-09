package routes

import (
	"Golang_Backend/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(Router *gin.Engine) {
	Router.GET("user", handlers.FindAll)
	Router.POST("user", handlers.CreateUser)
	Router.DELETE("user/:id", handlers.DeleteUserById)
}
