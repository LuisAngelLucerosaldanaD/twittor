package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion es la funcion que te lista la relacion de usuaios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := MongoCN.Database("twiter_clone")
	coll := db.Collection("relacion")

	condicion := bson.M{"usuarioid": t.UsuarioID, "usuariorelacionid": t.UsuarioRelacionID}
	var resultado models.Relacion
	fmt.Println(resultado)
	err := coll.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
