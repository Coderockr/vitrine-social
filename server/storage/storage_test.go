package storage_test

import (
	"log"
	"os"
	"testing"

	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {
	err := godotenv.Load("../config/test.env")
	os.Setenv("STORAGE_LOCAL_PATH", os.TempDir())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = storage.Connect()
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}

	os.Setenv("STORAGE_LOCAL_PATH_FRONTEND", os.TempDir())
	_, err = storage.ConnectFrontend()
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
}
