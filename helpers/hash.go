package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashedPassword(password string) (string, error) {
	if password == "" || password == "undefined" {
		return "", errors.New("password can not be null")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyHashedPassword(hashed_password string, password string) error {
	if password == "" || password == "undefined" {
		return errors.New("password can not be null")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err
}
