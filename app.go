package main

import (
	"Golang_Backend/db"
	middlewares "Golang_Backend/middleware"
	"Golang_Backend/routes"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	server := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not Load environment variables: %w", err)
	}

	db.ConnectToMongo()
	defer db.MongoDB.Disconnect(context.TODO())

	authRoutes := server.Group("/auth")
	routes.RegisterAuthRoutes(authRoutes)

	userRoutes := server.Group("/user")
	userRoutes.Use(middlewares.RequireAuth())
	routes.RegisterUserRoutes(userRoutes)

	todoRoutes := server.Group("/todo")
	todoRoutes.Use(middlewares.RequireAuth())
	routes.RegisterTodoRoutes(todoRoutes)

	server.Run(":3000")

}
