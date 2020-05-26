package bd

import (
	"context"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la parada final con la bd para unsertar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	coll := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)
	result, err := coll.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
