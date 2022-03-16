package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Fatalf("Error:%s", message)
	c.AbortWithStatusJSON(statusCode, error{Message: message})
	return
}
