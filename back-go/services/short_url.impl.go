package services

import (
	"context"
	"log"
	"time"

	"github.com/jbalibrea1/short-url-back/back-go/models"
	"github.com/jbalibrea1/short-url-back/back-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShortUrlServiceImpl struct {
	Collection *mongo.Collection
	ctx        context.Context
}

func NewShortUrlService(coll *mongo.Collection, ctx context.Context) ShortUrlService {
	return &ShortUrlServiceImpl{coll, ctx}
}

func (s *ShortUrlServiceImpl) GetAll() ([]*models.ShortUrl, error) {
	// Crear una variable que apunte a un slice de ShortUrl
	var allShortsURLs []*models.ShortUrl

	// Buscar todos los documentos en la colección y almacenarlos en el cursor
	//(cursor es un iterador que permite recorrer los documentos)
	cursor, err := s.Collection.Find(s.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(s.ctx)

	// Iterar sobre el cursor y decodificar los documentos en la estructura entryURL
	for cursor.Next(s.ctx) {
		entryURL := &models.ShortUrl{}
		if err := cursor.Decode(&entryURL); err != nil {
			log.Println("Failed to decode document:", err)
			continue
		}
		allShortsURLs = append(allShortsURLs, entryURL) // Añadir el doc decodificado al slice
	}

	// Responder con el slice
	return allShortsURLs, nil
}

func (s *ShortUrlServiceImpl) GetSingle(shortURL string) (*models.ShortUrl, error) {
	// 1) Crear variable para almacenar el documento decodificado
	var result *models.ShortUrl

	// 2) Definir filtro y actualización
	filter := bson.M{"shortURL": shortURL}             // Filtro para buscar el documento con la shortURL proporcionada
	update := bson.M{"$inc": bson.M{"totalClicks": 1}} // Actualizar el campo totalClicks por cada clic

	// 3) Definir opciones para la búsqueda y actualización del documento
	after := options.After
	upsert := false
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,  // Devolver el documento después de la actualización
		Upsert:         &upsert, // No crear un nuevo documento si no se encuentra
	}

	// 4) Buscar, actualizar el documento con la shortURL proporcionada y decodificarlo en la variable result
	res := s.Collection.FindOneAndUpdate(s.ctx, filter, update, &opt)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, res.Err()
	}

	// 5) Decodificar el documento res en la variable result
	res.Decode(&result)
	return result, nil
}

func (s *ShortUrlServiceImpl) Create(shortURL *models.OnlyURL) (*models.ShortUrl, error) {
	//TODO QUIZÁS CAMBIAR EN LA FUNCIÓN
	// 1) Validar la URL proporcionada y generar un shortID único
	parsedURL, err := utils.IsValidURL(shortURL.URL)
	if err != nil {
		return nil, err
	}

	sid, err := utils.GenerateUniqueShortID(s.Collection)
	if err != nil {
		return nil, err
	}

	// 2) Crear un documento con la URL, shortURL y otros campos
	result := &models.ShortUrl{
		URL:         parsedURL,
		ShortURL:    sid,
		TotalClicks: 0,
		CreatedAt:   time.Now(),
	}

	meta, err := utils.GetMetadata(result.URL)
	if err != nil {
		log.Println("Failed to get metadata:", err)
	} else {
		result.Title = &meta.Title
		result.Description = &meta.Description
		result.Logo = &meta.Favicon
	}

	res, err := s.Collection.InsertOne(s.ctx, result)
	if err != nil {
		return nil, err
	}
	result.ID = res.InsertedID.(primitive.ObjectID)

	return result, nil
}

func (s *ShortUrlServiceImpl) GetRedirect(shortURL string) (*models.OnlyURL, error) {
	// 1) Crear variable para almacenar el documento decodificado
	var result *models.ShortUrl
	var ret *models.OnlyURL
	// 2) Definir filtro y actualización
	filter := bson.M{"shortURL": shortURL}             // Filtro para buscar el documento con la shortURL proporcionada
	update := bson.M{"$inc": bson.M{"totalClicks": 1}} // Actualizar el campo totalClicks por cada clic

	// 3) Definir opciones para la búsqueda y actualización del documento
	after := options.After
	upsert := false
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,  // Devolver el documento después de la actualización
		Upsert:         &upsert, // No crear un nuevo documento si no se encuentra
	}

	// 4) Buscar, actualizar el documento con la shortURL proporcionada y decodificarlo en la variable result
	res := s.Collection.FindOneAndUpdate(s.ctx, filter, update, &opt)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, res.Err()
	}

	// 5) Decodificar el documento res en la variable result
	res.Decode(&result)

	// 6) Responder con la URL original
	ret = &models.OnlyURL{URL: result.URL}
	return ret, nil
}
