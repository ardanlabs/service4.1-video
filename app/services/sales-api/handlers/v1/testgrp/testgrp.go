package testgrp

import (
	"encoding/json"
	"net/http"
)

// Test is our example route.
func Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
