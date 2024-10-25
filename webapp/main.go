package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var openSearchURLs []string

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Message   string `json:"message"`
}

func main() {
	// Load OpenSearch URLs from the environment variable
	openSearchURLs = strings.Split(os.Getenv("OPENSEARCH_URL"), ",")

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

// logHandler handles incoming HTTP requests and logs them to a random OpenSearch node
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

	// Send the log entry to a random OpenSearch node
	nodeURL := openSearchURLs[rand.Intn(len(openSearchURLs))]
	resp, err := http.Post(fmt.Sprintf("%s/logs/_doc", nodeURL), "application/json", bytes.NewBuffer(logData))
	if err != nil {
		http.Error(w, "Error sending log to OpenSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Respond back to the user
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged request: %s %s\n", r.Method, r.URL.Path)
}
