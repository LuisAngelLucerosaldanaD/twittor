package routers

import (
	"net/http"

	"github.com/LuisAngelLucerosaldanaD/bd"
	"github.com/LuisAngelLucerosaldanaD/models"
)

/*AltaRelacion es la ruta que me permite insertar la relacion*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "El parametro ID es necesario", http.StatusBadRequest)
		return
	}
	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = id

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar unsertar relacion", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
