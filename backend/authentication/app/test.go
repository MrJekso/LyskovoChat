package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ для проверки подписи токена
var secretKey = []byte("my-secret-key")

// Функция для верификации и извлечения данных из JWT
func verifyToken(tokenString string) (*jwt.Token, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		// Возвращаем секретный ключ для проверки подписи
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем, валиден ли токен
	if !token.Valid {
		return nil, fmt.Errorf("токен невалиден")
	}

	return token, nil
}

func main() {
	// Пример Bearer Token (JWT)
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzkyMDEwOTcsInVzZXJuYW1lIjoiZW5vdCJ9.NngO93eS7CKkiFWEkdxKNteVWYTXsNFKPvrcKhJBWzE"

	// Верифицируем токен и извлекаем данные
	token, err := verifyToken(tokenString)
	if err != nil {
		log.Fatalf("Ошибка при верификации токена: %v", err)
	}

	// Извлекаем claims (утверждения) из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("Данные из токена:")
		fmt.Println("Username:", claims["username"])
		fmt.Println("Expiration:", claims["exp"])
	} else {
		log.Fatalf("Не удалось извлечь claims из токена")
	}
}