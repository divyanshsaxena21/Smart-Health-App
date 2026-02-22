package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"smart-health-app/handlers"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Handle Google Cloud credentials from environment variable (for Render deployment)
	setupGoogleCredentials()

	// Set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/process-image", handlers.ProcessImage)
	mux.HandleFunc("/health", healthCheck)

	// Configure CORS: allow localhost, any vercel.app subdomain, and any onrender.com subdomain
	corsHandler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			if origin == "" {
				return false
			}
			// Allow local development
			if strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "http://127.0.0.1") || strings.HasPrefix(origin, "https://localhost") {
				return true
			}
			// Allow any Vercel subdomain
			if strings.HasSuffix(origin, ".vercel.app") {
				return true
			}
			// Allow Render/app hostname subdomains
			if strings.HasSuffix(origin, ".onrender.com") {
				return true
			}
			return false
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := corsHandler.Handler(mux)

	// Get port from environment or default to 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// setupGoogleCredentials writes GCP credentials from env var to a temp file
// This is needed for cloud deployments where you can't upload files
func setupGoogleCredentials() {
	// Check if credentials JSON is provided as environment variable
	credentialsJSON := os.Getenv("GOOGLE_CREDENTIALS_JSON")
	if credentialsJSON == "" {
		log.Println("GOOGLE_CREDENTIALS_JSON not set, using GOOGLE_APPLICATION_CREDENTIALS path")
		return
	}

	// Write credentials to a temp file
	tmpFile := "/tmp/google-credentials.json"
	err := os.WriteFile(tmpFile, []byte(credentialsJSON), 0600)
	if err != nil {
		log.Fatalf("Failed to write credentials file: %v", err)
	}

	// Set the environment variable to point to the temp file
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", tmpFile)
	log.Println("Google Cloud credentials configured from environment variable")
}

// healthCheck handles the /health endpoint for Render health checks
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
