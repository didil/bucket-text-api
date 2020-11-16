package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StartServer starts the server
func StartServer() error {
	// try load env if .env file found
	err := LoadEnv(".env")
	if err != nil {
		// skip file not found errors to allow .env file to be optional
		if err.Error() != fmt.Sprintf("open .env: no such file or directory") {
			return fmt.Errorf("load env err: %v", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Initializing gcp service ...\n")
	gcpSVC, err := NewGCPService()
	if err != nil {
		return err
	}

	app := &App{
		GCPSvc: gcpSVC,
	}

	log.Printf("Initializing router ...\n")
	mux := BuildRouter(app)

	fmt.Printf("Listening on port %s\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
