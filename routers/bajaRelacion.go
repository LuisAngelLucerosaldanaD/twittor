package routers

import (
	"net/http"

	"github.com/LuisAngelLucerosaldanaD/bd"
	"github.com/LuisAngelLucerosaldanaD/models"
)

/*BajaRelacion es la funcion que me permite borrar una relacion*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = id
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar relacion", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
