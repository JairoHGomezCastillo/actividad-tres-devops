package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jaigomez/actividad-tres-devops/internal/models"
)

// Lista hardcodeada (en memoria)
var razas = []models.Raza{
	{ID: 1, Nombre: "French Poodle", Cualidades: "Inteligente, sociable, elegante"},
	{ID: 2, Nombre: "Bulldog Francés", Cualidades: "Cariñoso, tranquilo, leal"},
}

// HandleRazas maneja los endpoints REST (GET y POST)
func HandleRazas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(razas)

	case http.MethodPost:
		var nuevaRaza models.Raza
		if err := json.NewDecoder(r.Body).Decode(&nuevaRaza); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		nuevaRaza.ID = len(razas) + 1
		razas = append(razas, nuevaRaza)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nuevaRaza)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// GetRazas devuelve la lista actual para otros handlers (como el HTML)
func GetRazas() []models.Raza {
	return razas
}
