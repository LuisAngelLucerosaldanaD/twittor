package models

/*Relacion es el modelo de las relaciones entre usuarios*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuarioRelacionID" json:"usuarioRelacionID"`
}
