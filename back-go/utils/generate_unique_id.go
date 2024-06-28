package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// GenerateUniqueShortID genera un ID corto único que no existe en la base de datos
func GenerateUniqueShortID(c *gin.Context) (string, error) {
	var sid string

	for {
		// Genera un nuevo ID corto
		var newSid string
		newSid, err := shortid.Generate()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short ID"})
			log.Println("Failed to generate short ID:", err)
			return "", err
		}
		sid = newSid

		// Verifica si el ID corto generado ya existe en la base de datos
		if !IsShortUrlExist(sid) {
			break
		}
	}

	return sid, nil
}
