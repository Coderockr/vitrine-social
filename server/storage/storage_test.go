package storage_test

import (
	"log"
	"os"
	"testing"

	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	env := os.Getenv("VITRINESOCIAL_ENV")
	err := godotenv.Load("../config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = storage.Connect()
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
}

func TestSaveFile(t *testing.T)   {}
func TestDeleteFile(t *testing.T) {}
