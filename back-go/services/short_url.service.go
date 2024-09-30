package services

import "github.com/jbalibrea1/short-url-back/back-go/models"

type ShortUrlService interface {
	GetAllShortURLs() ([]*models.ShortUrl, error)
	GetSingleShortURL(string) (*models.ShortUrl, error)
	CreateShortURL(*models.CreateShortURL) (*models.ShortUrl, error)
}
