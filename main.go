package main

import (
	"log"

	"github.com/LuisAngelLucerosaldanaD/bd"
	"github.com/LuisAngelLucerosaldanaD/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}
	handlers.Manejadores()
}
