package status

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes describes routes for the status features
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/health", healthHandler)
	router.Get("/readiness", readinessHandler)
	return router
}

// Health godoc
// @Summary Return API health
// @Success 200
// @Router /status/health [get]
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Readiness godoc
// @Summary Return API readiness
// @Success 200
// @Router /status/readiness [get]
func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
