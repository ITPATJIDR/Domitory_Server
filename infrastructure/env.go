package infrastructure

import (
	"log"

	"github.com/joho/godotenv"
)

func InitENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Please consider environment variables: %s", err)
	}
}
