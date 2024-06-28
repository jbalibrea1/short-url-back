package utils

import (
	"context"
	"short-url/database"
	"short-url/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// IsShortUrlExist verifica si el shortURL ya existe en la base de datos
func IsShortUrlExist(shortURL string) bool {
	collection, _ := database.GetCollection("myDBwGO", "shorturls")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result models.ShortUrl
	err := collection.FindOne(ctx, bson.M{"short_URL": shortURL}).Decode(&result)
	return err == nil
	// if err != nil {
	// 	return false
	// }
	// return true
}
