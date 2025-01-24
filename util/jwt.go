package util

import (
	jwt2 "github.com/golang-jwt/jwt/v4"
	"security-go/response"
	"time"
)

func GenerateJWT(authResponse response.AuthResponse) (string, error) {
	// Crear los claims con los datos del usuario y menús
	claims := jwt2.MapClaims{
		"user":      authResponse.User,
		"menus":     authResponse.Menus,
		"entidadId": authResponse.EntidadId,
		"exp":       time.Now().Add(24 * time.Hour).Unix(), // Expiración en 24 horas
		"iss":       GetServiceName(),                      // Nombre de la aplicación
	}

	token := jwt2.NewWithClaims(jwt2.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(GetJWT()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
