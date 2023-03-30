package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcsvinicius/go-api-jobs/schemas"
)

// @BasePath /api/v1

// @Summary Get List Openings
// @Description Get list of a job opening
// @Tags Openings
// @Produce json
// @Success 200 {object} GetOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
