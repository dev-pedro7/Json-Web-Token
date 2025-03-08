package main

import (
	"fmt"
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

func ValidadateJWT(tokenString string) (*jwt.Token, error) {
	// Fazer o parsing do token e validar
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar se o algoritmo de assinatura está correto
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// 1: Faz o parsing do token para extrair os dados e validar a assinatura.
// 2: Verifica se o método de assinatura é do tipo HMAC SHA256.
// 3: Se tudo estiver certo, retorna o token validado.
// 4: Caso ocorra um erro (exemplo: token inválido ou expirado), ele retorna um erro
