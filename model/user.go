package model

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"

	"github.com/mitchellh/mapstructure"
)

// User : Usuario del sistema
type User struct {
	ID       primitive.ObjectID `json:"id"              bson:"_id"`
	Email    string             `json:"email"           bson:"email"`
	Name     string             `json:"name"          bson:"name"`
	Rol      string             `json:"rol"             bson:"rol"`
	Hash     string             `json:"_hash"           bson:"_hash,omitempty"`
	Password string             `json:"password,omitempty"           bson:"password,omitempty"`
}

// LoadFromContext : Traer usuario desde contexto
func (userModel *User) LoadFromContext(c *gin.Context) *User {
	claims := jwt.ExtractClaims(c)
	var user User
	err := mapstructure.Decode(claims["user"], &user)
	if err != nil {
		panic(err)
	}
	user.ID, err = primitive.ObjectIDFromHex(claims["user"].(map[string]interface{})["id"].(string))
	user.Hash = ""
	return &user
}
