package main

import (
	"context"
	"log"
	"os"

	"github.com/jbalibrea1/short-url-back/back-go/controllers"
	"github.com/jbalibrea1/short-url-back/back-go/database"
	"github.com/jbalibrea1/short-url-back/back-go/middleware"
	"github.com/jbalibrea1/short-url-back/back-go/router"
	"github.com/jbalibrea1/short-url-back/back-go/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	mongoClient *database.MongoClient
	// collection  *mongo.Collection
	ctx                     context.Context
	shortUrlService         services.ShortUrlService
	shortUrlController      controllers.ShortUrlController
	shortURLRouteController router.ShortUrlRouterController

	r *gin.Engine
)

func init() {
	log.Println("Starting server...")

	// 1) Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2) Crear un contexto
	ctx = context.Background()

	// 3) Conectar a la base de datos
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}
	mongoClient, err = database.NewMongoClient(mongodbURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	log.Println("Connected to MongoDB!")

	// 4) Obtener la colección de la base de datos
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME not set in environment")
	}
	dbCollection := os.Getenv("DB_COLLECTION")
	if dbCollection == "" {
		log.Fatal("DB_COLLECTION not set in environment")
	}
	collection, err := mongoClient.GetCollection(dbName, dbCollection)
	if err != nil {
		log.Fatal("Error getting shorturl collection:", err)
	}

	// 5) Crear instancias de servicios y controladores
	shortUrlService = services.NewShortUrlService(collection, ctx)
	shortUrlController = controllers.NewShortUrlController(shortUrlService)
	shortURLRouteController = router.NewShortUrlRoutes(shortUrlController)

	// 6) Crear una instancia de Gin
	r = gin.Default()
}

func main() {
	// 1) Cerrar la conexión a la base de datos al finalizar el programa
	defer mongoClient.Close(ctx)

	// 2) Configurar middlewares
	r.Use(middleware.CorsMiddleware())
	r.NoRoute(middleware.UnknownEndpoint)

	router := r.Group("/api")
	// 3) Configurar rutas de la API / añadimos a la instancia de Gin las rutas de la API
	shortURLRouteController.SetupRoutes(router)

	// 5) Iniciar el servidor Gin
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
