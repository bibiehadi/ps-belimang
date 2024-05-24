package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt, _ := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hashedPassword), err
}

func CompareHashAndPassword(hasedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
	return err == nil
}
