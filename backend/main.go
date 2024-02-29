package main

import "net/http"

func main() {
	// Create new mux
	mux := http.NewServeMux()

	// Add handlers

	// Start server
	if err := http.ListenAndServe("localhost:3001", mux); err != nil {
		panic(err)
	}
}
