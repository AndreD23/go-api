package router

import (
	docs "github.com/AndreD23/go-api/docs"
	"github.com/AndreD23/go-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group(basePath)
	{
		v1.GET("/opportunity", handler.ListOpportunitiesHandler)
		v1.GET("/opportunity/:id", handler.ShowOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpeningHandler)
		v1.PUT("/opportunity/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opportunity/:id", handler.DeleteOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
