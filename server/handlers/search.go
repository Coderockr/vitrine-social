package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Coderockr/vitrine-social/server/db/repo"
)

type (
	// SearchNeedRepository represet operations for need repository.
	SearchNeedRepository interface {
		Search(text string, categoriesID []int, organizationsID int64, page int64) ([]repo.DBSearch, error)
	}
)

// SearchHandler search needs for the term
func SearchHandler(sR SearchNeedRepository) func(w http.ResponseWriter, r *http.Request) {
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

		IDsplited := strings.Split(queryValues.Get("categories"), ",")
		categoriesID := make([]int, len(IDsplited))
		for i, v := range IDsplited {
			categoriesID[i], err = strconv.Atoi(v)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", v))
			}
		}

		needs, err := sR.Search(queryValues.Get("text"), categoriesID, orgID, page)

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, convertDBToNeed(needs))
	}
}

func convertDBToNeed(dbSearch []repo.DBSearch) []searchResultJSON {
	var need []searchResultJSON
	need = make([]searchResultJSON, len(dbSearch))
	for i, s := range dbSearch {
		need[i] = searchResultJSON{
			ID:               s.ID,
			Title:            s.Title,
			Description:      s.Description,
			RequiredQuantity: s.RequiredQuantity,
			ReachedQuantity:  s.ReachedQuantity,
			Unity:            s.Unity,
			DueDate:          s.DueDate,
			Category: categoryJSON{
				ID:   s.CategoryID,
				Name: s.CategoryName,
				Icon: s.CategoryIcon,
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
