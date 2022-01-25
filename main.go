package main

import (
	"log"
	"net/http"
	"time"

	"groovin.dev/inv-api/pkg/state"
	"groovin.dev/inv-api/pkg/utils"
)

// Main Function
func main() {
	// Create a new state manager
	s := state.StateManager{}

	// Initialize the state manager
	s.Init()

	// Register the routes
	r := utils.RegisterRouter(&s)

	// Start the server
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Print that we are starting the server
	log.Println("Starting server on port 8000")

	log.Fatal(srv.ListenAndServe())
}
