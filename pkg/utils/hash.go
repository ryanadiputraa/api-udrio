package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passowrd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passowrd), 14)
	return string(bytes), err
}

func CheckPasswordWithHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
