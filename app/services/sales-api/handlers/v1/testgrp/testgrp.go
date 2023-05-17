package testgrp

import (
	"context"
	"encoding/json"
	"net/http"
)

// Test is our example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Validate the data
	// Call into the business layer
	// Return errors
	// Handle OK response

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(status)
}
