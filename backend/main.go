package main

import (
	"log"
	"net/http"
	"os"

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

	// Set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/process-image", handlers.ProcessImage)

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"https://cal-qulate.vercel.app",
			"https://calqulate.ayushsharma.site",
		},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
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
