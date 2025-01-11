package routes

import "github.com/gin-gonic/gin"

func RegisterTodoRoutes(router *gin.RouterGroup) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Stuff"})
	})

}
