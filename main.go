package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a Chi router
	r := chi.NewRouter()

	// Add middleware (logger and recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status":  "healthy",
			"service": "seek-clarity",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// API v1 routes
	r.Route("/api/v1", func(r chi.Router) {
		// Welcome endpoint
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response := map[string]string{
				"message": "Welcome to Seek Clarity API",
				"version": "1.0.0",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		})

		// Example CRUD endpoints placeholder
		r.Get("/users", getUsers)
		r.Post("/users", createUser)
		r.Get("/users/{id}", getUserByID)
		r.Put("/users/{id}", updateUser)
		r.Delete("/users/{id}", deleteUser)
	})

	// Start server on port 8080
	log.Println("Starting Seek Clarity server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Placeholder handler functions
func getUsers(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Get all users",
		"data":    []interface{}{},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "User created successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response := map[string]string{
		"message": "Get user by ID",
		"id":      id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response := map[string]string{
		"message": "User updated successfully",
		"id":      id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	response := map[string]string{
		"message": "User deleted successfully",
		"id":      id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
