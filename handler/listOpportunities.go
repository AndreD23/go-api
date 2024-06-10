package handler

import (
	"github.com/AndreD23/go-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
