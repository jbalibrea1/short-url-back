package main

import (
	"log"
	"os"
	"short-url/database"
	"short-url/router"

	"github.com/joho/godotenv"
)

func main() {
	// 1) Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongodbURI := os.Getenv("MONGODB_URI")
	if mongodbURI == "" {
		log.Fatal("MONGODB_URI not set in environment")
	}

	// 2) Conectar a la base de datos
	if err := database.Connect(mongodbURI); err != nil {
		log.Fatal(err)
	}

	// 3) Configurar el router, rutas y middlewares
	r := router.SetupRouter()

	// 4) Iniciar el servidor
	log.Fatal(r.Run(":8080"))
}
