package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// SearchNeedRepository represet operations for need repository.
	SearchNeedRepository interface {
		Search(text string, categoriesID []int, organizationsID int64, orderBy string, order string, page int64) ([]model.SearchNeed, error)
	}
)

// SearchHandler search needs for the term
func SearchHandler(sR SearchNeedRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()

		var orgID int64
		var categoriesID []int
		var err error

		if len(queryValues.Get("page")) < 1 {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Parametros inválidos"))
			return
		}

		if len(queryValues.Get("org")) > 0 {
			orgID, err = strconv.ParseInt(queryValues.Get("org"), 10, 64)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender a organização: %s", queryValues.Get("org")))
				return
			}
		}

		if len(queryValues.Get("categories")) > 0 {
			idSplited := strings.Split(queryValues.Get("categories"), ",")
			categoriesID = make([]int, len(idSplited))
			for i, v := range idSplited {
				categoriesID[i], err = strconv.Atoi(v)
				if err != nil {
					HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender a categoria: %s", v))
					return
				}
			}
		}

		page, err := strconv.ParseInt(queryValues.Get("page"), 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("page")))
		}

		text := queryValues.Get("text")
		orderBy := queryValues.Get("orderBy")
		order := queryValues.Get("order")

		needs, err := sR.Search(text, categoriesID, orgID, orderBy, order, page)

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, convertDBToNeed(needs))
	}
}

func convertDBToNeed(searchNeed []model.SearchNeed) []searchResultJSON {
	var need []searchResultJSON
	need = make([]searchResultJSON, len(searchNeed))
	for i, s := range searchNeed {
		need[i] = searchResultJSON{
			ID:               s.ID,
			Title:            s.Title,
			Description:      s.Description,
			RequiredQuantity: s.RequiredQuantity,
			ReachedQuantity:  s.ReachedQuantity,
			Unit:             s.Unity,
			DueDate:          s.DueDate,
			CreatedAt:        s.CreatedAt,
			UpdatedAt:        s.UpdatedAt,
			Category: categoryJSON{
				ID:   s.CategoryID,
				Name: s.CategoryName,
				Slug: s.CategorySlug,
			},
			Organization: baseOrganizationJSON{
				ID:   s.OrganizationID,
				Name: s.OrganizationName,
				Logo: s.OrganizationLogo,
				Slug: s.OrganizationSlug,
			},
		}
	}

	return need
}
