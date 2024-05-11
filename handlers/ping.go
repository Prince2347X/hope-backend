package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}