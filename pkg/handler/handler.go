package handler

import (
	"github.com/404th/todolist/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// AUTH
	r.POST("/auth/sign-up", h.signUp)
	r.POST("/auth/sign-in", h.signIn)

	// API
	r.POST("/api/lists/", h.createList)
	r.GET("/api/lists/", h.getAllLists)
	r.GET("/api/lists/:id", h.getListByID)
	r.PUT("/api/lists/:id", h.updateList)
	r.DELETE("/api/lists/:id", h.deleteList)

	// ITEMS
	r.POST("/api/lists/:id/items/", h.createItem)
	r.GET("/api/lists/:id/items/", h.getAllItems)
	r.GET("/api/lists/:id/items/:item_id", h.getItemByID)
	r.PUT("/api/lists/:id/items/:item_id", h.updateItem)
	r.DELETE("/api/lists/:id/items/:item_id", h.deleteItem)

	return r
}
