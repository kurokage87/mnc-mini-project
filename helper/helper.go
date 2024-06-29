package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

func GenerateToken(userID uuid.UUID, duration time.Duration) (token string, err error) {
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	})

	token, err = tokenJwt.SignedString([]byte("your_secret_key"))
	if err != nil {
		return
	}

	return
}
