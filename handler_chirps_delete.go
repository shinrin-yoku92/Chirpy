package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/shinrin_yoku92/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid or missing token", err)
		return
	}

	userID, err := auth.ValidateJWT(tok, cfg.secretKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid or missing token", err)
		return
	}

	// Verify that the chirp belongs to the user
	idStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(idStr)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "You do not have permission to delete this chirp", nil)
		return
	}

	// Proceed to delete the chirp
	cfg.db.DeleteChirpByID(r.Context(), chirpID)
	respondWithJSON(w, http.StatusNoContent, nil)
}
