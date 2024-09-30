package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnknownEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
}
