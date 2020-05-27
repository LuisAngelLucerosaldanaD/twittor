package models

/*Tweet captura el mensaje que viene del body*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
