package handler

import (
	"net/http"

	"github.com/gerfalcao/go-ahead.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error;err!=nil{
		sendError(ctx, http.StatusInternalServerError, "error listening openings")
	}

	sendSuccess(ctx, "list-openings", openings)

}