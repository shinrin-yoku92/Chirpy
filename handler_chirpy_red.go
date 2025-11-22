package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/shinrin_yoku92/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeChirpyRed(w http.ResponseWriter, r *http.Request) {
	type webhookRequest struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	decoder := json.NewDecoder(r.Body)
	req := webhookRequest{}
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if req.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "unsupported event type", nil)
		return
	}

	_, err = auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid API key", err)
		return
	}

	_, err = cfg.db.UpgradeChirpyRed(r.Context(), req.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "user not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "failed to upgrade user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
