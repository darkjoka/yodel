package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateToken(id uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id.String(),
	})
	return token.SignedString(getSecret())

}
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["id"].(string)
		return username, nil
	}
	return "", err
}

func getSecret() []byte {
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		panic("No secret found")
	}
	return []byte(secret)

}
