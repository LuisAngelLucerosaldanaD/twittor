package bd

import (
	"context"
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los tweets de mis seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := MongoCN.Database("twiter_clone")
	coll := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})

	condiciones = append(condiciones, bson.M{
		"&lookup": bson.M{
			"from":         "tweet",
			"lacalField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})
	cursor, err := coll.Aggregate(ctx, condiciones)
	var results []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}
	return results, true
}
