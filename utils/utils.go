package utils

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func HttpErrorHandler(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}

func GetDateString() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

func ErrorHandler(message string) error {
	return errors.New(message)
}