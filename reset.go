package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev platform"))
		return
	}

	cfg.fileserverHits.Store(0)

	err := cfg.db.Reset(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset database" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database reset successfully"))
}
