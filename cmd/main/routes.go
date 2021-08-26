package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Returns a multiplexer
func Routes() *chi.Mux {
	// Create a new multiplexer
	mux := chi.NewMux()

	// Take advantage of middlewares provided by go.chi/chi.
	// These are going to be applied to each route we create.
	// They are "global"
	mux.Use(
		middleware.Logger,    // Log every http request
		middleware.Recoverer, // Recover if panic occurs
	)

	// Handlers
	mux.Get("/api/notes", nil)          // Get all notes
	mux.Get("/api/notes/<id:int>", nil) // Get a single note
	mux.Post("/api/notes", nil)         // Create a note
	mux.Put("/api/notes/<id:int>", nil) // Modify a note

	return mux
}
