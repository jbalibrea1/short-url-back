package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ShortUrl es la estructura que representa un documento en la colección de MongoDB
type ShortUrl struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	URL         string             `bson:"url" json:"url"`
	Title       *string            `bson:"title,omitempty" json:"title,omitempty"`
	Logo        *string            `bson:"logo,omitempty" json:"logo,omitempty"`
	Description *string            `bson:"description,omitempty" json:"description,omitempty"`
	ShortURL    string             `bson:"shortURL" json:"shortURL"`
	TotalClicks int                `bson:"totalClicks" json:"totalClicks"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}

type CreateShortURL struct {
	URL string `json:"url" binding:"required"`
}
