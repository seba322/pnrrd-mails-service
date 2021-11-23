package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hierarchy struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Lat        string             `json:"lat" bson:"lat,omitempty"`
	Lng        string             `json:"lng" bson:"lng,omitempty"`
	Provincias []interface{}      `json:"provincias" bson:"provincias,omitempty"`
}

func (hierarchyModel *Hierarchy) FindAll() ([]Hierarchy, error) {
	col, _, ctx := GetCollection(CollectionNameHierarchy)

	hierarchies := []Hierarchy{}

	cursor, err := col.Find(ctx, bson.D{{}})

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &hierarchies)
	return hierarchies, err

}
