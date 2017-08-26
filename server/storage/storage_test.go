package storage_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Coderockr/vitrine-social/server/storage"
	"github.com/joho/godotenv"
)

func TestDial(t *testing.T) {
	env := os.Getenv("CODENATION_ENV")
	err := godotenv.Load("../config/" + env + ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	location, err := storage.Dial()
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
	defer location.Close()
}

func TestContainer(t *testing.T) {
	profile := "eminetto"
	location, _ := storage.Dial()
	_, err := storage.Container(location, profile)
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
	location.Close()
	location, _ = storage.Dial()
	container, err := storage.Container(location, profile)
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
	defer location.Close()
	name := profile + ".txt"

	content := []byte("temporary file's content")
	tmpfile, _ := ioutil.TempFile("", "example")
	_, _ = tmpfile.Write(content)

	defer os.Remove(tmpfile.Name()) // clean up

	resume, _ := os.Open(tmpfile.Name())
	var buff bytes.Buffer
	fileSize, _ := buff.ReadFrom(resume)
	_, _ = resume.Seek(0, 0)
	_, err = container.Put(name, resume, fileSize, nil)
	if err != nil {
		t.Errorf("expected %v result %s", nil, err)
	}
	dat, _ := ioutil.ReadFile("/tmp/eminetto/eminetto.txt")
	if string(dat) != string(content) {
		t.Errorf("expected %s result %s", string(content), string(dat))
	}
}
