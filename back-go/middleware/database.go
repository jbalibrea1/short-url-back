package middleware

import (
	"log"
	"net/http"
	"short-url/database"

	"github.com/gin-gonic/gin"
)

// DatabaseMiddleware es un middleware que obtiene la colección de la base de datos y la añade al contexto Gin
func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection, err := database.GetCollection("myDBwGO", "shorturls")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
			log.Println("Failed to connect to database:", err)
			c.Abort()
			return
		}

		c.Set("collection", collection)
		c.Next()
	}
}
