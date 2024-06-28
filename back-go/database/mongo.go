package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(uri string) error {
	var err error

	// 1) Contexto con timeout para la conexión a la base de datos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2) Crear la conexión a MongoDB
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	// 3) Verificar la conexión
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// 4) Asignar el cliente a la variable global
	Client = client

	log.Println("Connected to MongoDB!")
	return nil
}

// GetCollection obtiene una colección de la base de datos
func GetCollection(database, collection string) (*mongo.Collection, error) {
	if Client == nil {
		return nil, fmt.Errorf("MongoDB client is nil")
	}

	coll := Client.Database(database).Collection(collection)
	if coll == nil {
		return nil, fmt.Errorf("collection %s not found in database %s", collection, database)
	}

	return coll, nil
}
