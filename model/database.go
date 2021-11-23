package model

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

// GetCollection : Obtener coleccion desde la bd
func GetCollection(collection string) (*mongo.Collection, *mongo.Client, context.Context) {
	s := MI.Client
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	return MI.DB.Collection(collection), s, ctx
}

// LoadDatabase : Carga la base de datos y devuelve la session correspondiente
func LoadDatabase() {
	credential := options.Credential{
		AuthSource: os.Getenv("DB_DB"),
		Username:   os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URL")).SetAuth(credential))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB_DB")),
	}

	CreateIndex(MI.DB)

}
