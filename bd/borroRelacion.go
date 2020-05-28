package bd

import (
	"context"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
)

/*BorroRelacion es la funcion que me borra la relacion en la base de datos*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := MongoCN.Database("twiter_clone")
	coll := db.Collection("relacion")

	_, err := coll.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
