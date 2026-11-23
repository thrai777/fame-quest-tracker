// Command server runs the fame-quest-tracker API server.
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thrai777/fame-quest-tracker/internal/domain"
	"github.com/thrai777/fame-quest-tracker/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           server.New(domain.SeedQuests()).Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("fame-quest-tracker listening on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
