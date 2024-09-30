package utils

import (
	"context"
	"time"

	"github.com/jbalibrea1/short-url-back/back-go/models"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GenerateUniqueShortID genera un ID corto único que no existe en la base de datos
func GenerateUniqueShortID(coll *mongo.Collection) (string, error) {
	var sid string

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
		// Genera un nuevo ID corto
		var newSid string
		newSid, err := shortid.Generate()
		if err != nil {
			return "", err
		}
		sid = newSid

		var result models.ShortUrl

		// Busca el ID corto en la base de datos
		// Si no se encuentra el ID corto en la base de datos, es único, se rompe el bucle
		if err = coll.FindOne(ctx, bson.M{"short_URL": sid}).Decode(&result); err != nil {
			break
		}
	}

	return sid, nil
}
