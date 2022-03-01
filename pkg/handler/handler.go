package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	// AUTHENTICATION
	auth := routes.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	// LISTS
	lists := routes.Group("/lists")
	{
		lists.GET("/")
		lists.GET("/:id")
		lists.POST("/")
		lists.PUT("/:id")
		lists.DELETE("/:id")
		lists.GET("/:id/items")
		lists.POST("/:id/items")
	}

	// ITEMS
	items := routes.Group("/items")
	{
		items.GET("/:item_id")
		items.PUT("/:item_id")
		items.DELETE("/:item_id")
	}

	return routes
}

// package handler

// import (
// 	"github.com/gin-gonic/gin"
// )

// type Handler struct {
// }

// func (h *Handler) InitRoutes() *gin.Engine {
// 	router := gin.New()

// 	auth := router.Group("/auth")
// 	{
// 		auth.POST("/sign-up")
// 		auth.POST("/sign-in")
// 	}

// 	api := router.Group("/api")
// 	{
// 		lists := api.Group("/lists")
// 		{
// 			lists.POST("/")
// 			lists.GET("/")
// 			lists.GET("/:id")
// 			lists.PUT("/:id")
// 			lists.DELETE("/:id")

// 			items := lists.Group(":id/items")
// 			{
// 				items.POST("/")
// 				items.GET("/")
// 				items.GET("/:item_id")
// 				items.PUT("/:item_id")
// 				items.DELETE("/:item_id")
// 			}
// 		}
// 	}

// 	return router
// }
