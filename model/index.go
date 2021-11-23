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
	CollectionNameUser      = "User"
	CollectionNameInventory = "NewInventory"
)

// CreateIndex : Crea indices en las colecciones de mongo
func CreateIndex(session *mongo.Database) {

	// Creacion de indices

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	InventoryIndex := []mongo.IndexModel{

		{
			Keys: bsonx.Doc{{Key: "institucion_id", Value: bsonx.Int32(1)}, {Key: "hierarchy_id", Value: bsonx.Int32(1)}},
		},
	}

	_, err := session.Collection(CollectionNameInventory).Indexes().CreateMany(context.Background(), InventoryIndex, opts)

	if err != nil {
		log.Printf("Error al crear indice %s en : %s, %s", "TaggedData", "idData", err)
		panic("No se pudo crear indice")
	}

	log.Print("Creando indices")
}
