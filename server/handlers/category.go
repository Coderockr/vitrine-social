package handlers

import (
	"fmt"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// CategoryRepository represet operations for category repository.
	CategoryRepository interface {
		GetAll() ([]model.Category, error)
	}
)

// GetAllCategoriesHandler will retrive the data from all categories
func GetAllCategoriesHandler(cR CategoryRepository, nR NeedRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		categories, err := cR.GetAll()
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível buscar as categorias: %s", err))
			return
		}

		cJSON := make([]categoryWithCountJSON, 0, len(categories))

		for _, c := range categories {
			cJSON = append(cJSON, categoryWithCountJSON{
				categoryJSON: categoryJSON{
					ID:   c.ID,
					Name: c.Name,
					Icon: c.Icon,
				},
				NeedsCount: c.NeedsCount,
			})
		}

		HandleHTTPSuccess(w, cJSON)
	}
}
