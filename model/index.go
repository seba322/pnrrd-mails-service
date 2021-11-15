package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// Nombre de la base de datos

// Constantes de nombres de colecciones
const (
	CollectionNameDog       = "Dog"
	CollectionNameUser      = "User"
	CollectionNameInventory = "NewInventory"
	CollectionNameForm      = "Form"
)

// CreateIndex : Crea indices en las colecciones de mongo
func CreateIndex(session *mongo.Database) {

	// Creacion de indices
	formIndex := []mongo.IndexModel{

		{
			Keys: bsonx.Doc{{Key: "_id", Value: bsonx.Int32(1)}},
		},
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	_, err := session.Collection(CollectionNameForm).Indexes().CreateMany(context.Background(), formIndex, opts)
	if err != nil {
		log.Printf("Error al crear indice %s en : %s, %s", "TaggedData", "idData", err)
		panic("No se pudo crear indice")
	}

	log.Print("Creando indices")
}
