package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisAngelLucerosaldanaD/bd"
)

/*LeoTweets leo los tweets devueltos de la base de datos*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe de enviar el parametro ID", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe de enviar el parametro Página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe de enviar el parametro Página con un valor mayor a 0 ", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correto := bd.LeoTweets(ID, pag)
	if correto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(respuesta)
}
