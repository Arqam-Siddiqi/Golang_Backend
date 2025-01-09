package routes

import (
	"Golang_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(Router *gin.Engine) {
	Router.GET("user", controllers.FindAll)
	Router.POST("user", controllers.CreateUser)
	Router.DELETE("user/:id", controllers.DeleteUserById)
	Router.GET("user/:id", controllers.FindById)
	Router.PUT("user/:id", controllers.UpdateUserById)
}
