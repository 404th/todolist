package handler

import (
	"github.com/gin-gonic/gin"
)

func NewHandler() *gin.Engine {
	r := gin.New()

	r.POST("/auth/sign-up", signUp)
	r.POST("/auth/sign-in", signIn)

	r.GET("/task", getAllTasks)
	r.GET("/task/:id", getTaskByID)
	r.POST("/task", createTask)
	r.PUT("/task/:id", updateTask)
	r.DELETE("/task/:id", deleteTask)

	return r
}
