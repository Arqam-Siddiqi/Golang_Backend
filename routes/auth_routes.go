package routes

import (
	"Golang_Backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

}
