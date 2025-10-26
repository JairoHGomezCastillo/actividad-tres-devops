package main

import (
	"fmt"
	"net/http"

	"github.com/jaigomez/actividad-tres-devops/internal/handlers"
	"github.com/jaigomez/actividad-tres-devops/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	// Endpoint REST (GET y POST)
	mux.HandleFunc("/razas", handlers.HandleRazas)

	// Página HTML
	mux.HandleFunc("/", handlers.RenderRazaTable)

	// Aplicar middleware de logging
	loggedMux := middleware.LoggingMiddleware(mux)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
