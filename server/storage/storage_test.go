package storage_test

import (
	"log"
	"testing"

	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	err := godotenv.Load("../config/test.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = storage.Connect()
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
}
