package middlewares

import (
	"Golang_Backend/db"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			log.Fatal("JWT_SECRET is not set in .env file.")
		}

		authorizationHeader := ctx.GetHeader("Authorization")
		if authorizationHeader == "" {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization Headers have not been set."})
			return
		}

		fullToken := strings.Split(authorizationHeader, " ")
		if len(fullToken) != 2 || strings.ToLower(fullToken[0]) != "bearer" || fullToken[1] == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(fullToken[1], func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		var claims jwt.MapClaims
		var ok bool
		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		_id := claims["_id"]
		objectId, err := primitive.ObjectIDFromHex(_id.(string))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		exists := db.UserModel.FindOne(context.TODO(), bson.M{"_id": objectId})

		if exists.Err() != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", objectId) // Example of setting user ID from claims

		ctx.Next()
	}

}
