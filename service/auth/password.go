package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(psswd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(psswd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
