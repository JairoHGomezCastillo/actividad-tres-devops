package handlers

import (
	"fmt"
	"net/http"
)

// RenderRazaTable renderiza una tabla HTML con las razas
func RenderRazaTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="es">
		<head>
			<meta charset="UTF-8">
			<title>Razas de Perros</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f8f9fa;
					padding: 40px;
				}
				h1 {
					text-align: center;
					color: #333;
				}
				table {
					width: 60%%;
					margin: 30px auto;
					border-collapse: collapse;
					box-shadow: 0 2px 8px rgba(0,0,0,0.1);
					background-color: #fff;
				}
				th, td {
					padding: 12px 15px;
					text-align: left;
					border-bottom: 1px solid #ddd;
				}
				th {
					background-color: #007BFF;
					color: white;
				}
				tr:hover {
					background-color: #f1f1f1;
				}
			</style>
		</head>
		<body>
			<h1>Tabla de Razas de Perros</h1>
			<table>
				<tr>
					<th>ID</th>
					<th>Nombre</th>
					<th>Cualidades</th>
				</tr>`)

	for _, r := range GetRazas() {
		fmt.Fprintf(w, `
				<tr>
					<td>%d</td>
					<td>%s</td>
					<td>%s</td>
				</tr>`, r.ID, r.Nombre, r.Cualidades)
	}

	fmt.Fprintf(w, `
			</table>
		</body>
		</html>
	`)
}
