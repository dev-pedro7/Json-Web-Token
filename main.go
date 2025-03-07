package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Chave secreta usada para assinar e validar o token
var jwtSecret = []byte("minhaChaveSecreta")

// Função para gerar um token JWT
func GenerateJWT() (string, error) {
	// Criar as claims (dados do token)
	claims := jwt.MapClaims{
		"username": "user123",
		"exp":      time.Now().Add(time.Minute * 2).Unix(), // Expira em 2 minutos
	}

	// Criar o token com algoritmo de assinatura HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assinar o token com a chave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
