package main

import (
	"fmt"
	"net/http"

	"github.com/jaigomez/actividad-tres-devops/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Endpoint REST (GET y POST)
	mux.HandleFunc("/razas", handlers.HandleRazas)

	// Página HTML
	mux.HandleFunc("/", handlers.RenderRazaTable)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
