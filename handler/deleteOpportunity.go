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

func DeleteOpeningHandler(ctx *gin.Context) {
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

	// Delete Opportunity
	if err := db.Delete(&opportunity); err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opportunity with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opportunity", opportunity)
}
