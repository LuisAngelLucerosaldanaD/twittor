package jwt

import (
	"time"

	"github.com/LuisAngelLucerosaldanaD/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT es la funcion que me devuelve el token de acceso*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("NodeSystem_job")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellido,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
