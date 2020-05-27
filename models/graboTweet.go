package models

import (
	"time"
)

/*GraboTweet es el modelo del tweet*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userId,omitempty`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty`
}
