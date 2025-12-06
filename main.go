package main

import (
	"log"
	"net/http"

	"veediver-backend/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a Chi router
	r := chi.NewRouter()

	// Add middleware (logger and recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup routes
	routes.SetupRoutes(r)

	// Start server on port 8080
	log.Println("Starting Veediver Backend server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
