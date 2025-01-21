package auth

import (
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func GenerateToken(message string) string {
	godotenv.Load(".env")

	secretKey := os.Getenv("SECRET_KEY")

	unixTime := time.Now().UnixNano()
	uniqueTimeString := strconv.Itoa(int(unixTime))

	hasher := sha512.New()

	hasher.Write([]byte(message))
	hasher.Write([]byte(uniqueTimeString))

	key := hasher.Sum([]byte(secretKey))

	return fmt.Sprintf("%x", key)
}
