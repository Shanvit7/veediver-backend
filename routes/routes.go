package routes

import (
	"veediver-backend/handlers"

	"github.com/go-chi/chi/v5"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *chi.Mux) {
	// Health check endpoint
	r.Get("/health", handlers.Health)
}
