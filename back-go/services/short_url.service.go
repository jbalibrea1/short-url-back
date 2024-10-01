package services

import "github.com/jbalibrea1/short-url-back/back-go/models"

type ShortUrlService interface {
	GetAll() ([]*models.ShortUrl, error)
	GetSingle(string) (*models.ShortUrl, error)
	Create(*models.OnlyURL) (*models.ShortUrl, error)
	GetRedirect(string) (*models.OnlyURL, error)
}
