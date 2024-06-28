package handlers

import (
	"log"
	"net/http"
	"short-url/models"
	"short-url/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetShortURL(c *gin.Context) {
	// 1) Crear contexto
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	// 2) Verifica si el parámetro shortURL está presente en la URL y no está vacío
	shortURL := c.Param("shortURL")
	if shortURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short URL not provided"})
		return
	}

	// 3) Obtener la colección de la base de datos
	coll, ok := utils.GetCollectionFromContext(c)
	if !ok {
		return // Se maneja el error dentro de GetCollectionFromContext
	}

	var result models.ShortUrl // Variable para almacenar el documento encontrado

	// 4) Definir filtro y actualización
	filter := bson.M{"shortURL": shortURL}             // Filtro para buscar el documento con la shortURL proporcionada
	update := bson.M{"$inc": bson.M{"totalClicks": 1}} // Actualizar el campo totalClicks por cada clic

	// 5) Definir opciones para la búsqueda y actualización del documento
	after := options.After
	upsert := false
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,  // Devolver el documento después de la actualización
		Upsert:         &upsert, // No crear un nuevo documento si no se encuentra
	}

	// 6) Buscar, actualizar el documento con la shortURL proporcionada y decodificarlo en la variable result
	res := coll.FindOneAndUpdate(ctx, filter, update, &opt)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve document"})
			log.Println("Failed to retrieve document:", res)
		}
		return
	}

	// 7) Decodificar el documento res en la variable result
	res.Decode(&result)

	// 8) Responder con el documento decodificado
	c.JSON(http.StatusOK, result)
}

func CreateShortURL(c *gin.Context) {
	// 1) Craer contexto
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	// 2) Intenta hacer el binding del JSON recibido en la estructura input
	var input struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format or body is empty"})
		return
	}

	// 3) Verificar si el campo URL está presente y no está vacío
	if input.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'url' field is required"})
		return
	}

	// 4) Obtener la colección de la base de datos
	coll, ok := utils.GetCollectionFromContext(c)
	if !ok {
		return // Se maneja el error dentro de GetCollectionFromContext
	}

	// 5) Generar un ID corto y único para la URL
	sid, err := utils.GenerateUniqueShortID(c)
	if err != nil {
		return
	}

	// 6) Configurar los campos restantes de input
	result := models.ShortUrl{
		ID:          primitive.NewObjectID(),
		URL:         utils.AddHTTPPrefixIfNeeded(input.URL),
		ShortURL:    sid,
		TotalClicks: 0,
		CreatedAt:   time.Now(),
	}

	// 7) Obtener metadatos de la URL proporcionada
	meta, err := utils.GetMetadata(result.URL)
	if err != nil {
		log.Println("Failed to get metadata:", err)
	} else {
		result.Title = &meta.Title
		result.Description = &meta.Description
		result.Logo = &meta.Favicon
	}
	log.Println("Metadata:", meta)

	// Generar el código QR
	qrCode, err := utils.GenerateQRCode(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate QR code"})
		return
	}
	result.QRCode = qrCode

	// 8) Insertar el documento en la base de datos
	_, err = coll.InsertOne(ctx, result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		log.Println("Failed to insert document:", err)
		return
	}

	// 9) Responder con el documento insertado
	c.JSON(http.StatusOK, result)
}

func GetShortURLs(c *gin.Context) {
	// 1) Crear contexto
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	// 2) Crear un slice
	var allShortsURLs []models.ShortUrl

	// 3) Obtener la colección de la base de datos
	coll, ok := utils.GetCollectionFromContext(c)
	if !ok {
		return // Se maneja el error dentro de GetCollectionFromContext
	}

	// 4) Buscar todos los documentos en la colección y almacenarlos en el cursor
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve documents"})
		log.Println("Failed to retrieve documents:", err)
		return
	}
	defer cursor.Close(ctx)

	// 5) Iterar sobre el cursor y decodificar los documentos en la estructura entryURL
	for cursor.Next(ctx) {
		var entryURL models.ShortUrl                    // Variable para almacenar cada single documento
		cursor.Decode(&entryURL)                        // Decodifica el documento en la variable entryURL
		allShortsURLs = append(allShortsURLs, entryURL) // Agrega la entrada al slice
	}

	// 6) Responder con el slice
	c.JSON(http.StatusOK, allShortsURLs)
}
