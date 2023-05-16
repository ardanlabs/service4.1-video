package testgrp

import (
	"encoding/json"
	"net/http"
)

// Test is our example route.
func Test(w http.ResponseWriter, r *http.Request) {

	// Validate the data
	// Call into the business layer
	// Return errors
	// Handle OK response

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
