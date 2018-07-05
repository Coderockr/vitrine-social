package testutil

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// NewFileUploadRequest creates a multipart/form-data request
func NewFileUploadRequest(path, method string, params map[string]string, files map[string]string) *http.Request {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for paramName, path := range files {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fi, err := file.Stat()
		if err != nil {
			panic(err)
		}
		file.Close()

		part, err := writer.CreateFormFile(paramName, fi.Name())
		if err != nil {
			panic(err)
		}
		part.Write(fileContents)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}

	r, _ := http.NewRequest(method, path, body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	return r
}
