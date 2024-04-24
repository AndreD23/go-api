package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opportunities",
	})
}
