package handlers

import (
	"errors"
	"net/http"

	"github.com/graymeta/stow"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	imageStorage interface {
		GetImage(*model.Token, int64) error
	}
)

// GetImageHandler search an image and return
func GetImageHandler(container stow.Container) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()
		item, err := container.Item(queryValues.Get("url"))
		if err != nil {
			HandleHTTPError(w, http.StatusNotFound, errors.New("Imagem não encontrada"))
		}

		reader, err := item.Open()
		if err != nil {
			HandleHTTPError(w, http.StatusNotFound, errors.New("Erro ao processar imagem"))
		}
		defer reader.Close()

		HandleHTTPSuccessImage(w, reader)
	}
}
