package main

import (
	"sync/atomic"

	"github.com/shinrin_yoku92/Chirpy/internal/db"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *db.Queries
	platform       string
}
