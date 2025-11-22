package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shinrin_yoku92/Chirpy/internal/auth"
	"github.com/shinrin_yoku92/Chirpy/internal/db"
)

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type loginRequest struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := loginRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	userRecord, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, userRecord.PasswordHash)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", err)
		return
	}

	accessToken, err := auth.MakeJWT(userRecord.ID, cfg.secretKey, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create JWT token", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	expires := time.Now().Add(60 * 24 * time.Hour) // 60 days
	_, err = cfg.db.InsertRefreshTokens(r.Context(), db.InsertRefreshTokensParams{
		Token:     refreshToken,
		UserID:    userRecord.ID,
		ExpiresAt: expires,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't store refresh token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:          userRecord.ID,
			CreatedAt:   userRecord.CreatedAt,
			UpdatedAt:   userRecord.UpdatedAt,
			Email:       userRecord.Email,
			IsChirpyRed: userRecord.IsChirpyRed,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
