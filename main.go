package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/celestiaorg/supply/supply"
)

func circulatingHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "`date` query parameter is required", http.StatusBadRequest)
		return
	}

	// Parse date
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	// Calculate circulating supply
	circulating := supply.CirculatingSupply(t)

	// Write response
	fmt.Fprintf(w, "Circulating Supply for %s: %d", dateStr, circulating)
}

func availableHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "`date` query parameter is required", http.StatusBadRequest)
		return
	}

	// Parse date
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD.", http.StatusBadRequest)
		return
	}

	// Calculate available supply
	available := supply.AvailableSupply(t)

	// Write response
	fmt.Fprintf(w, "Available Supply for %s: %d", dateStr, available)
}

func main() {
	http.HandleFunc("/circulating", circulatingHandler)
	http.HandleFunc("/available", availableHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
