package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRenderRazaTable(t *testing.T) {
	// Crear un request simulado
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// Ejecutar el handler
	RenderRazaTable(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Validar el código de estado
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("esperaba status 200, obtuve %d", resp.StatusCode)
	}

	// Leer el contenido HTML generado
	body := w.Body.String()

	// Verificar que tenga partes clave del HTML
	if !strings.Contains(body, "<h1>Tabla de Razas de Perros</h1>") {
		t.Error("no se encontró el título esperado en el HTML")
	}

	// Verificar que aparezcan las razas conocidas
	if !strings.Contains(body, "French Poodle") {
		t.Error("no se encontró la raza 'French Poodle' en el HTML")
	}
	if !strings.Contains(body, "Bulldog Francés") {
		t.Error("no se encontró la raza 'Bulldog Francés' en el HTML")
	}

	// Verificar que se esté generando una tabla
	if !strings.Contains(body, "<table>") {
		t.Error("no se encontró la tabla en el HTML generado")
	}
}
