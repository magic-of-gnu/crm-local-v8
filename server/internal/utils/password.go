package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, hash_cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hash_cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
