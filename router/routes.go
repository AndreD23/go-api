package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunity/id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET opportunities",
			})
		})
		v1.GET("/opportunity", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET opportunity",
			})
		})
		v1.POST("/opportunity", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "POST opportunity",
			})
		})
		v1.PUT("/opportunity/id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT opportunity",
			})
		})
		v1.DELETE("/opportunity/id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE opportunity",
			})
		})
	}
}
