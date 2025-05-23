package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(hashedUserPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedUserPassword), []byte(password))
	return err == nil
}