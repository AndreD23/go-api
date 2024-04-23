package router

import "github.com/gin-gonic/gin"

func BuildAndRun() {
	router := gin.Default()
	loadRoutes(router)
	router.Run(":3001")
}
