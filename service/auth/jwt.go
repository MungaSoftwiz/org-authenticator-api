package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
