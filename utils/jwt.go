package utils

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(GetJWTSecret())
	return t
}

func GetJWTSecret() []byte {
	jwtSecretString := os.Getenv("JWT_SECRET")
	return []byte(jwtSecretString)
}
