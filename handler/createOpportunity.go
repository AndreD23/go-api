package handler

import (
	"fmt"
	"github.com/AndreD23/go-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// CreateOpeningHandler
// @Summary CreateOpeningHandler
// @Description Create a new job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param request body CreateOpportunityRequest true "Request body"
// @Success 200 {object} CreateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [post]
func CreateOpeningHandler(ctx *gin.Context) {
	// Difines the fields to get from json body
	request := CreateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating create opportunity request: %v", err.Error())
		sendError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	fmt.Println(request.Remote)
	fmt.Println(*request.Remote)

	opportunity := schemas.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opportunity", opportunity)
}
