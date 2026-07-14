package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    "healthy",
			"message":   "Golang demo corriendo en K8s!",
			"version":   "1.0.0",
			"timestamp": time.Now(),
		})
	})
	http.ListenAndServe(":8080", nil)
}