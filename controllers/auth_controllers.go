package controllers

import (
	"Golang_Backend/models"
	"Golang_Backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var token string
	var err error
	token, err = utils.CreateJwt(user.ID.(primitive.ObjectID).Hex())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"email": user.Email,
	})

}

func Login(ctx *gin.Context) {

	var loginRequest models.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	var user *models.User
	var err error
	user, err = models.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	var token string
	token, err = utils.CreateJwt(user.ID.(primitive.ObjectID).Hex())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"email": user.Email,
	})

}
