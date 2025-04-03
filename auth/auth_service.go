package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedUserPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedUserPassword))
	return err == nil
}