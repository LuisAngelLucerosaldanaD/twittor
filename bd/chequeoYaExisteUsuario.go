package bd

import (
	"context"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email como parametro y verifica si ya esta en la base de datos*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twiter_clone")
	coll := db.Collection("usuarios")
	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := coll.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
