package handler

import (
	"errors"
	"fmt"
	"github.com/AndreD23/go-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// @BasePath /api/v1

// UpdateOpeningHandler
// @Summary UpdateOpeningHandler
// @Description Update a job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id query string true "Opportunity identification"
// @Param request body UpdateOpportunityRequest true "Request body"
// @Success 200 {object} UpdateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity/id [patch]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramPath").Error())
		return
	}

	opportunity := schemas.Opportunity{}

	intID, err := strconv.Atoi(id)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "paramPath").Error())
		return
	}

	// Find Opportunity
	err = db.First(&opportunity, intID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	fmt.Println(opportunity)

	// Update Opportunity
	request := UpdateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating update opportunity request: %v", err.Error())
		sendError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Update Opportunity fields
	if request.Role != "" {
		opportunity.Role = request.Role
	}
	if request.Company != "" {
		opportunity.Company = request.Company
	}
	if request.Location != "" {
		opportunity.Location = request.Location
	}
	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}
	if request.Link != "" {
		opportunity.Link = request.Link
	}
	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	// Save the updated Opportunity
	err = db.Save(&opportunity).Error
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to update opportunity: %v", err.Error()))
		return
	}

	sendSuccess(ctx, "update-opportunity", opportunity)
}
