package handler

import (
	"github.com/AndreD23/go-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// ListOpportunitiesHandler
// @Summary ListOpportunitiesHandler
// @Description List all jobs opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [get]
func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
