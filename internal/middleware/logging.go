package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware registra cada petición con método, ruta y duración.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("← %s %s completado en %v", r.Method, r.URL.Path, duration)
	})
}
