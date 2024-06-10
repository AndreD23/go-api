package handler

import (
	"fmt"
	"github.com/AndreD23/go-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type DeleteOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type ShowOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type ListOpportunityResponse struct {
	Message string                        `json:"message"`
	Data    []schemas.OpportunityResponse `json:"data"`
}

type UpdateOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}
