package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	secret := "test_secret"
	userID := uuid.New()
	expiresIn := time.Minute * 5

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}

	returnedUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	if returnedUserID != userID {
		t.Fatalf("Expected userID %v, got %v", userID, returnedUserID)
	}
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	secret := "test_secret"
	invalidToken := "invalid.token.string"

	_, err := ValidateJWT(invalidToken, secret)
	if err == nil {
		t.Fatalf("Expected error for invalid token, got nil")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	secret := "test_secret"
	wrongSecret := "wrong_secret"
	userID := uuid.New()
	expiresIn := time.Minute * 5

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}

	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Fatalf("Expected error for wrong secret, got nil")
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	secret := "test_secret"
	userID := uuid.New()
	expiresIn := -time.Minute * 5 // Token already expired

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("Failed to make JWT: %v", err)
	}

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatalf("Expected error for expired token, got nil")
	}
}

func TestValidateJWT_InvalidClaims(t *testing.T) {
	secret := "test_secret"
	invalidToken := "eyJhbGciOi" + // Truncated invalid token
		".eyJpc3MiOiJjaGlycHkiLCJpYXQiOjE2MDAwMDAwMDAsImV4cCI6MTYwMDAwMDAwMCwic3ViIjoiSW52YWxpZFVVSUQifQ" +
		".invalidsignature"

	_, err := ValidateJWT(invalidToken, secret)
	if err == nil {
		t.Fatalf("Expected error for invalid claims, got nil")
	}
}
