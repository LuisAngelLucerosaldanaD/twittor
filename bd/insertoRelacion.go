package bd

import (
	"context"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
)

/*InsertoRelacion es la funcion que me permite insertar los datos de la relacion a la bd */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := MongoCN.Database("twiter_clone")
	coll := db.Collection("relacion")
	_, err := coll.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
