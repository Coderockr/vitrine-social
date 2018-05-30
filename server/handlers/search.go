package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// SearchNeedRepository represet operations for need repository.
	SearchNeedRepository interface {
		Search(text string, categoriesID []int, organizationsID int64, page int64) ([]model.Need, error)
	}
)

// SearchHandler search needs for the term
func SearchHandler(nR SearchNeedRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()

		if len(queryValues) < 1 {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Parametros inválidos"))
			return
		}

		orgID, err := strconv.ParseInt(queryValues.Get("org"), 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("org")))
		}

		page, err := strconv.ParseInt(queryValues.Get("page"), 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("page")))
		}

		needs, err := nR.Search(queryValues.Get("text"), []int{2, 3, 5, 7, 11, 13}, orgID, page)

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		nJSON := make([]needJSON, 0, len(needs))

		var dueDate *jsonTime
		for _, n := range needs {
			if n.DueDate != nil {
				dueDate = &jsonTime{*n.DueDate}
			}

			nJSON = append(nJSON, needJSON{
				ID:               n.ID,
				Title:            n.Title,
				Description:      n.Description,
				RequiredQuantity: n.RequiredQuantity,
				ReachedQuantity:  n.ReachedQuantity,
				Unity:            n.Unity,
				DueDate:          dueDate,
				Category: categoryJSON{
					ID:   n.Category.ID,
					Name: n.Category.Name,
					Icon: n.Category.Icon,
				},
				Organization: baseOrganizationJSON{
					ID:   n.Organization.ID,
					Name: n.Organization.Name,
					Logo: n.Organization.Logo,
					Slug: n.Organization.Slug,
				},
				Images: needImagesToImageJSON(n.Images),
				Status: string(n.Status),
			})
		}

		HandleHTTPSuccess(w, nJSON)
	}
}
