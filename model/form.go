package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Form struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreationDate time.Time          `json:"creationDate" bson:"creation_date,omitempty"`
	ModifiedDate time.Time          `json:"modifiedDate" bson:"modified_date,omitempty"`
	TypeForm     string             `json:"typeForm" bson:"type_form,omitempty"`
	Hierarchy    string             `json:"hierarchy" bson:"hierarchy,omitempty"`
	Sections     []interface{}      `json:"sections" bson:"sections,omitempty"`
}

func (formModel *Form) Create(formDoc *Form) error {

	col, _, ctx := GetCollection(CollectionNameForm)
	formDoc.ID = primitive.NewObjectID()
	_, err := col.InsertOne(ctx, formDoc)

	return err
}

func (formModel *Form) GetForm(typeForm string, hierarchy string) (*Form, error) {
	col, _, ctx := GetCollection(CollectionNameForm)

	var form Form
	query := bson.D{{"type_form", typeForm}, {"hierarchy", hierarchy}}

	err := col.FindOne(ctx, query).Decode(&form)

	return &form, err

}
