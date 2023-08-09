package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var account Account
	errBind := c.ShouldBindJSON(&account)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error validation, username and password is required",
		})
		return
	}

	if account.Username != "admin" || account.Password != "password" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token": tokenString,
	}) 
}
