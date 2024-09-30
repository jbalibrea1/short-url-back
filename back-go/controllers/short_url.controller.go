package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jbalibrea1/short-url-back/back-go/models"
	"github.com/jbalibrea1/short-url-back/back-go/services"
)

type ShortUrlController struct {
	shortUrlService services.ShortUrlService
}

func NewShortUrlController(service services.ShortUrlService) ShortUrlController {
	return ShortUrlController{service}
}

// GET /api/shorturl
func (suc *ShortUrlController) GetAllShortURLs(ctx *gin.Context) {
	// * El servicio esta encapsulado en el controlador
	// 1) Llamar al servicio para obtener un slice
	allShortsURLs, err := suc.shortUrlService.GetAllShortURLs()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve documents"})
		log.Println("Failed to retrieve documents:", err)
		return
	}

	// 2) Responder con el slice obtenido
	ctx.JSON(http.StatusOK, allShortsURLs)
}

// GET /api/shorturl/:shortURL
func (suc *ShortUrlController) GetShortURL(ctx *gin.Context) {
	shortUrl := ctx.Param("shortURL")
	shortURL, err := suc.shortUrlService.GetSingleShortURL(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve short URL", "message": "Short URL not found " + shortUrl})
		return
	}
	ctx.JSON(http.StatusOK, shortURL)
}

// POST /api/shorturl -d '{"url": "https://www.google.com"}'
func (suc *ShortUrlController) CreateShortURL(ctx *gin.Context) {
	// Iniciamos la estructura para almacenar la URL
	var bodyURL *models.CreateShortURL

	// Decodificamos el JSON en la estructura bodyURL
	if err := ctx.ShouldBindJSON(&bodyURL); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "message": err.Error()})
		return
	}

	// Llamamos al servicio para crear la short URL
	shortURL, err := suc.shortUrlService.CreateShortURL(bodyURL)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Failed to retrieve short URL", "message": err.Error()})
		return
	}

	// Respondemos con la short URL creada
	ctx.JSON(http.StatusOK, shortURL)
}
