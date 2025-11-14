package auth

import (
	"errors"

	"github.com/alexedwards/argon2id"
)

func CheckPasswordHash(password, hash string) error {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("password does not match")
	}
	return nil
}
