package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't list chirps", err)
		return
	}

	var respChirps []Chirp
	for _, c := range chirps {
		respChirps = append(respChirps, Chirp{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body:      c.Body,
			UserID:    c.UserID,
		})
	}

	log.Printf("type=%T len=%d", respChirps, len(respChirps))
	b, _ := json.Marshal(respChirps)
	log.Printf("json=%s", b)

	respondWithJSON(w, http.StatusOK, respChirps)
}
