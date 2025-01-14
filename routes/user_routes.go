package routes

import (
	"Golang_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(Router *gin.RouterGroup) {
	Router.GET("/", controllers.FindAll)
	Router.GET("/self", controllers.GetUserByJwt)
	Router.DELETE("/", controllers.DeleteUserByJwt)
	Router.GET("/:id", controllers.FindById)
	Router.PUT("/:id", controllers.UpdateUserById)
}
