package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const openSearchURL = "http://opensearch-node:9200/logs/_doc" // OpenSearch endpoint

// LogEntry represents the log structure
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Message   string `json:"message"`
}

func main() {
	// Define a simple HTTP handler
	http.HandleFunc("/", logHandler)

	// Get the port from environment variables or default to 8080
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// logHandler handles incoming HTTP requests and logs them to OpenSearch
func logHandler(w http.ResponseWriter, r *http.Request) {
	logEntry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Method:    r.Method,
		Path:      r.URL.Path,
		Message:   "Received HTTP request",
	}

	// Convert log entry to JSON
	logData, err := json.Marshal(logEntry)
	if err != nil {
		http.Error(w, "Error creating log entry", http.StatusInternalServerError)
		return
	}

	// Send the log entry to OpenSearch
	resp, err := http.Post(openSearchURL, "application/json", bytes.NewBuffer(logData))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending log to OpenSearch: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Respond back to the user
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged request: %s %s\n", r.Method, r.URL.Path)
}
