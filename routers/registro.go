package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisAngelLucerosaldanaD/bd"
	"github.com/LuisAngelLucerosaldanaD/models"
)

/*Registro es la funcion para crear en la bd el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	if len(t.Password) < 8 {
		http.Error(w, "Debe de especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Yaexiste  un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Nose ha logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
