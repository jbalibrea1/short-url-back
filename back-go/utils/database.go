package utils

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollectionFromContext obtiene la colección de la base de datos del contexto Gin
func GetCollectionFromContext(c *gin.Context) (*mongo.Collection, bool) {
	collection, exists := c.Get("collection")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get database collection"})
		log.Println("Failed to get database collection from context")
		return nil, false
	}

	coll, ok := collection.(*mongo.Collection)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast database collection"})
		log.Println("Failed to cast database collection")
		return nil, false
	}

	return coll, true
}

// GetContextWithTimeout devuelve un contexto con timeout de 10 segundos
func GetContextWithTimeout() (context.Context, context.CancelFunc) {
	exp := 10 * time.Second
	return context.WithTimeout(context.Background(), exp)
}
