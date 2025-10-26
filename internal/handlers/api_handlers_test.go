package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaigomez/actividad-tres-devops/internal/models"
)

func TestHandleRazas_Get(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/razas", nil)
	w := httptest.NewRecorder()

	HandleRazas(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperaba status 200, obtuve %d", resp.StatusCode)
	}

	var got []models.Raza
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("error al decodificar respuesta: %v", err)
	}

	if len(got) == 0 {
		t.Errorf("se esperaba al menos una raza, pero no llegó ninguna")
	}
}

func TestHandleRazas_Post(t *testing.T) {
	nueva := models.Raza{
		Nombre:     "Labrador",
		Cualidades: "Fiel, activo, amigable",
	}
	body, _ := json.Marshal(nueva)

	req := httptest.NewRequest(http.MethodPost, "/razas", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	HandleRazas(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("esperaba status 201, obtuve %d", resp.StatusCode)
	}

	var got models.Raza
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("error al decodificar respuesta: %v", err)
	}

	if got.Nombre != nueva.Nombre {
		t.Errorf("nombre incorrecto: esperado %s, obtenido %s", nueva.Nombre, got.Nombre)
	}
}
