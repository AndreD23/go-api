package router

import (
	"github.com/AndreD23/go-api/handler"
	"github.com/gin-gonic/gin"
)

func loadRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunity", handler.ListOpportunitiesHandler)
		v1.GET("/opportunity/:id", handler.ShowOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpeningHandler)
		v1.PUT("/opportunity/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opportunity/:id", handler.DeleteOpeningHandler)
	}
}
