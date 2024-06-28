package router

import (
	"net/http"
	"short-url/handlers"
	"short-url/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 1) Crear un nuevo enrutador
	r := gin.Default()

	// 2) Obtener la colección de la base de datos y añadirla al contexto Gin
	r.Use(middleware.DatabaseMiddleware())

	// 3) setup cors
	// r.Use(corsMiddleware())
	config := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	r.Use(cors.New(config))

	// 4) Configurar rutas y controladores
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/shorturl", handlers.GetShortURLs)
	r.POST("/shorturl", handlers.CreateShortURL)
	r.GET("/shorturl/:shortURL", handlers.GetShortURL)

	return r
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}
// 		c.Next()
// 	}
// }
