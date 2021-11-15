package model

import (
	// "github.com/globalsign/mgo/bson"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Dog : Perro de un usuario
type Dog struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Age   int                `json:"age" bson:"age,omitempty"`
	Owner primitive.ObjectID `json:"owner" bson:"owner,omitempty"`
}

// Page : Pagina de resultado
type Page struct {
	Metadata []map[string]int `json:"metadata" bson:"metadata,omitempty"`
	Data     []interface{}    `json:"data" bson:"data,omitempty"`
}

// Create : Crear perro por ID
func (dogModel *Dog) Create(dogDoc *Dog) error {
	col, _ := GetCollection(CollectionNameDog)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dogDoc.ID = primitive.NewObjectID()
	_, err := col.InsertOne(ctx, dogDoc)

	return err
}

// Get : Obtener perro por ID
func (dogModel *Dog) Get(id string) (*Dog, error) {
	col, _ := GetCollection(CollectionNameDog)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var dogDoc Dog
	bsonId, err := primitive.ObjectIDFromHex(id)
	err = col.FindOne(ctx, bson.M{"_id": bsonId}).Decode(&dogDoc)

	return &dogDoc, err
}

// Update : Actualizar perro por ID
func (dogModel *Dog) Update(id string, dogDoc Dog) error {

	col, _ := GetCollection(CollectionNameDog)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bsonId, err := primitive.ObjectIDFromHex(id)
	_, err = col.UpdateByID(ctx, bson.M{"_id": bsonId}, bson.M{"$set": dogDoc})
	return err
}

// Delete : Eliminar perro por ID
func (dogModel *Dog) Delete(id string) error {

	col, _ := GetCollection(CollectionNameDog)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// defer session.Close()
	bsonId, err := primitive.ObjectIDFromHex(id)

	_, err = col.DeleteOne(ctx, bson.M{"_id": bsonId})
	return err
}

// Find : Obtener perro
func (dogModel *Dog) Find(query bson.M) ([]Dog, error) {

	col, _ := GetCollection(CollectionNameDog)
	// defer session.Close()
	dogs := []Dog{}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := col.Find(ctx, query)
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &dogs)
	return dogs, err
}

// FindPaginate : Obtener perro
func (dogModel *Dog) FindPaginate(query bson.D, limit int, offset int) (Page, error) {

	col, _ := GetCollection(CollectionNameDog)
	// defer session.Close()
	pag := []bson.M{{"$skip": offset}}
	if limit > 0 {
		pag = append(pag, bson.M{"$limit": limit})
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	matchStage := bson.D{{"$match", query}}
	countStage := bson.D{{"$facet", bson.D{{"metadata", bson.D{{"$count", "total"}}}, {"data", pag}}}}

	pageDoc := Page{}
	cursor, err := col.Aggregate(ctx, mongo.Pipeline{matchStage, countStage})
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &pageDoc)

	return pageDoc, err
}
