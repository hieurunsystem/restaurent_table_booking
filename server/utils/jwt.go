package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "kiemlam"

func GenarateToken(userId int64, gmail, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"gmail":  gmail,
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
