package main

import (
	"Golang_Backend/db"
	"Golang_Backend/routes"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	db.ConnectToMongo()
	defer db.MongoDB.Disconnect(context.TODO())

	routes.RegisterUserRoutes(server)

	server.Run(":3000")

}
