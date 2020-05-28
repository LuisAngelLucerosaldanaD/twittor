package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisAngelLucerosaldanaD/bd"
)

/*LeoTweetSeguidores es la ruta para poder leer los tweets de los seguidores*/
func LeoTweetSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe de enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Debe de enviar el parametro pagina como un entero mayor a 0", http.StatusBadRequest)
		return
	}
	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
