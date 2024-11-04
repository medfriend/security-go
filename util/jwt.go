package util

import (
	"github.com/dgrijalva/jwt-go"
	"security-go/response"
	"time"
)

func GenerateJWT(authResponse response.AuthResponse, secretKey string) (string, error) {
	// Crear los claims con los datos del usuario y menús
	claims := jwt.MapClaims{
		"user":  authResponse.User,
		"menus": authResponse.Menus,
		"exp":   time.Now().Add(24 * time.Hour).Unix(), // Expiración en 24 horas
		"iss":   "your-app-name",                       // Nombre de la aplicación
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
