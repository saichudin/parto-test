package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ShowErrMessage(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
}