package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcsvinicius/go-api-jobs/schemas"
)

func GetAllOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
