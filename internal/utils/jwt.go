// internal/utils/jwt.go
package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var jwtSecret = []byte("supersecretkey") // ⚠️ лучше из ENV

func SetJwtSecret() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func GetJwtSercret() []byte {
	return jwtSecret
}

func GenerateJWT(id uint) (string, error) {
	claims := jwt.MapClaims{
		"courier_id": id,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
