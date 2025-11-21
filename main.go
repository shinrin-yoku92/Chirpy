package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shinrin_yoku92/Chirpy/internal/db"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *db.Queries
	platform       string
	secretKey      string
}

func main() {
	const port = "8080"
	const filepathroot = "."

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM environment variable is not set")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()
	dbQueries := db.New(dbConn)

	cfg := &apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
		secretKey:      secretKey,
	}

	mux := http.NewServeMux()

	mux.Handle("/app/", cfg.middlewareMetrics(
		http.StripPrefix(
			"/app/", http.FileServer(
				http.Dir(
					filepathroot),
			),
		),
	),
	)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)

	mux.HandleFunc("POST /api/users", cfg.handlerCreateUser)
	mux.HandleFunc("PUT /api/users", cfg.handlerUpdateUserLogins)
	mux.HandleFunc("POST /api/login", cfg.handlerUserLogin)

	mux.HandleFunc("POST /api/chirps", cfg.handlerCreateChirp)
	mux.HandleFunc("GET /api/chirps", cfg.handlerListChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handlerGetChirpByID)

	mux.HandleFunc("POST /api/refresh", cfg.handlerTokenRefresh)
	mux.HandleFunc("POST /api/revoke", cfg.handlerRevoke)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on port %s", port)
	log.Printf("Serving files from %s", filepathroot)
	log.Fatal(srv.ListenAndServe())
}
