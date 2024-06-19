package jwt

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type JWTConfig struct {
	SecretKey string
}

func LoadJWTConfig() *JWTConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET"),
	}
}
