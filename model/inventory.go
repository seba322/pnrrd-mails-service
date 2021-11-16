package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Inventory: inventario para diferentes jerarquias
type Inventory struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	InstitucionId primitive.ObjectID `json:"institucionId" bson:"institucion_id"`
	PronvinciaId  primitive.ObjectID `json:"provinciaId" bson:"provincia_id,omitempty"`
	ComunaId      primitive.ObjectID `json:"comunaId" bson:"comuna_id,omitempty"`
	RegionId      primitive.ObjectID `json:"regionId" bson:"region_id,omitempty"`
	Hierarchy     string             `json:"hierarchy" bson:"hierarchy"`
	CreationDate  time.Time          `json:"creationDate" bson:"creation_date"`
	ModifiedDate  time.Time          `json:"modifiedDate" bson:"modified_date"`
	State         string             `json:"state" bson:"state,omitempty"`
	Details       interface{}        `json:"details" bson:"details,omitempty"`
}

const (
	PROVINCIAL_HIERARCHY = "PROVINCIAL"
	COMUNAL_HIERARCHY    = "COMUNAL"
	REGIONAL_HIERARCHY   = "REGIONAL"
	NACIONAL_HIERARCHY   = "NACIONAL"
)

var HierarchyMap = map[string]string{
	PROVINCIAL_HIERARCHY: "provincia_id",
	COMUNAL_HIERARCHY:    "comuna_id",
	REGIONAL_HIERARCHY:   "region_id",
}

func (inventoryModel *Inventory) UpsertManyInventory(inventories []Inventory) error {
	col, _, ctx := GetCollection(CollectionNameInventory)

	var operations []mongo.WriteModel

	for _, inv := range inventories {
		if inv.ID == primitive.NilObjectID {
			inv.ID = primitive.NewObjectID()
			op := mongo.NewInsertOneModel().
				SetDocument(inv)
			operations = append(operations, op)
		} else {
			op := mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": inv.ID}).
				SetUpdate(inv)
			op.SetUpsert(true)
			operations = append(operations, op)

		}
	}

	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)

	_, err := col.BulkWrite(ctx, operations, &bulkOption)

	return err
}

func (inventoryModel *Inventory) GetInventory(institutionId string, hierarchy string, hierarchyId string) ([]Inventory, error) {
	col, _, ctx := GetCollection(CollectionNameInventory)

	invs := []Inventory{}

	instId, err := primitive.ObjectIDFromHex(institutionId)
	hierarId, err := primitive.ObjectIDFromHex(hierarchyId)

	if err != nil {
		return nil, err
	}

	var query bson.D
	if hierarchy == NACIONAL_HIERARCHY {

		query = bson.D{{"institucion_id", instId}, {"hierarchy", NACIONAL_HIERARCHY}}
	} else {
		filterField := HierarchyMap[hierarchy]
		query = bson.D{{"institucion_id", instId}, {filterField, hierarId}}
	}

	cursor, err := col.Find(ctx, query)

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &invs)
	return invs, err
}
