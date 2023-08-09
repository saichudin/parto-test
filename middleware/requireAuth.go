package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Token")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" : "Token required",
		})
	}

	token, _:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Token expired",
			})
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message" : "Token invalid",
		})
	}
	
}