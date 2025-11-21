package main

import (
	"encoding/json"
	"net/http"

	"github.com/shinrin_yoku92/Chirpy/internal/auth"
	"github.com/shinrin_yoku92/Chirpy/internal/db"
)

func (cfg *apiConfig) handlerUpdateUserLogins(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

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

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	updatedUser, err := cfg.db.UpdateUserLogins(r.Context(), db.UpdateUserLoginsParams{
		Email:        params.Email,
		PasswordHash: hashedPassword,
		ID:           userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user logins", err)
		return
	}

	resp := response{
		User: User{
			ID:        updatedUser.ID,
			Email:     updatedUser.Email,
			UpdatedAt: updatedUser.UpdatedAt,
		},
	}
	respondWithJSON(w, http.StatusOK, resp)
}
