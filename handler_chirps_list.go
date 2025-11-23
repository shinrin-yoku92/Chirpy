package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {
	authID := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")

	if authID != "" {
		userID, err := uuid.Parse(authID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author_id", err)
			return
		}

		chirps, err := cfg.db.ListChirpsByUserID(r.Context(), userID)
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

		if sortOrder == "desc" {
			sort.Slice(respChirps, func(i, j int) bool {
				return respChirps[i].CreatedAt.After(respChirps[j].CreatedAt)
			})
		}

		respondWithJSON(w, http.StatusOK, respChirps)
	} else {
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

		if sortOrder == "desc" {
			sort.Slice(respChirps, func(i, j int) bool {
				return respChirps[i].CreatedAt.After(respChirps[j].CreatedAt)
			})
		}

		respondWithJSON(w, http.StatusOK, respChirps)
	}
}
