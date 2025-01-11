package routes

import (
	"Golang_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(Router *gin.RouterGroup) {
	Router.GET("/", controllers.FindAll)
	Router.POST("/", controllers.CreateUser)
	Router.DELETE("/:id", controllers.DeleteUserById)
	Router.GET("/:id", controllers.FindById)
	Router.PUT("/:id", controllers.UpdateUserById)
}
