package models

// Raza representa una raza de perro
type Raza struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Cualidades string `json:"cualidades"`
}
