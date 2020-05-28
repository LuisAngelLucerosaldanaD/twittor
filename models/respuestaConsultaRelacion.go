package models

/*RespuestaConsultaRelacion tiene el true o el false que se obtiene de consultar la relacion entre dos usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
