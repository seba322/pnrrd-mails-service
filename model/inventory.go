package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Inventory: inventario para diferentes jerarquias
type Inventory struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	InstitucionId bson.ObjectId `json:"institucionId" bson:"institucion_id"`
	PronvinciaId  bson.ObjectId `json:"provinciaId" bson:"provincia_id,omitempty"`
	ComunaId      bson.ObjectId `json:"ComunaId" bson:"comuna_id,omitempty"`
	CreationDate  time.Time     `json:"creationDate" bson:"creation_date"`
	ModifiedDate  time.Time     `json:"modifiedDate" bson:"modified_date"`
	State         string        `json:"state" bson:"state,omitempty"`
	Details       []interface{} `json:"details" bson:"details,omitempty"`
}
