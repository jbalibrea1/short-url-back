package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jbalibrea1/short-url-back/back-go/controllers"
)

type ShortUrlRouterController struct {
	shortUrlController controllers.ShortUrlController
}

func NewShortUrlRoutes(sucController controllers.ShortUrlController) ShortUrlRouterController {
	return ShortUrlRouterController{sucController}
}

func (suc *ShortUrlRouterController) SetupRoutes(r *gin.RouterGroup) {
	// El controlador esta encapsulado en el router
	//* gin pasa el contexto automáticamente a los controladores
	//* por lo que no es necesario pasarlo manualmente
	//* suc.shortUrlController.<controlador>(ctx)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/shorturl", suc.shortUrlController.GetAllShortURLs)
	r.GET("/shorturl/:shortURL", suc.shortUrlController.GetShortURL)
	r.POST("/shorturl", suc.shortUrlController.CreateShortURL)
}
