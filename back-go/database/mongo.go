package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

var mongoClientInstance *MongoClient

// NewMongoClient crea una instancia única de MongoClient
func NewMongoClient(uri string) (*MongoClient, error) {
	// 1) Si ya existe una instancia de MongoClient, devolverla
	if mongoClientInstance != nil {
		return mongoClientInstance, nil
	}

	// 2) Si no existe una instancia de MongoClient, crear una nueva
	mongoClientInstance = &MongoClient{}
	err := mongoClientInstance.Connect(uri)
	if err != nil {
		return nil, err
	}

	return mongoClientInstance, nil
}

// Función privada para conectar a MongoDB
func (mc *MongoClient) Connect(uri string) error {
	// 1) Crear un contexto con un timeout de 10 segundos, si no se conecta en ese tiempo, cancelar la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2) Conectar a MongoDB
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("error pinging MongoDB: %v", err)
	}

	mc.client = client

	return nil
}

// GetCollection obtiene una colección de la base de datos
func (mc *MongoClient) GetCollection(database, collection string) (*mongo.Collection, error) {
	if mc.client == nil {
		return nil, fmt.Errorf("MongoDB client is nil")
	}

	coll := mc.client.Database(database).Collection(collection)
	if coll == nil {
		return nil, fmt.Errorf("collection %s not found in database %s", collection, database)
	}

	return coll, nil
}

// Close cierra la conexión a MongoDB
func (mc *MongoClient) Close(ctx context.Context) {
	if mc.client != nil {
		mc.client.Disconnect(ctx)
		log.Println("Disconnected from MongoDB")
	}
}
