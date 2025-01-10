package auth

import (
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GenerateToken(message string) string {
	godotenv.Load(".env")

	secretKey := os.Getenv("SECRET_KEY")

	hasher := sha512.New()

	hasher.Write([]byte(message))

	key := hasher.Sum([]byte(secretKey))

	return fmt.Sprintf("%x", key)
}
