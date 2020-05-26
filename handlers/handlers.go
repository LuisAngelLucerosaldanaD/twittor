package handlers

import (
	"log"
	"net/http"
	"os"

	middleware "github.com/LuisAngelLucerosaldanaD/middlew"
	"github.com/LuisAngelLucerosaldanaD/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto y pongo a escuchar el server*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}