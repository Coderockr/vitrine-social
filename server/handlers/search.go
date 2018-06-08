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
		Search(text string, categoriesID []int, organizationsID int64, orderBy []string, page int64) ([]model.SearchNeed, error)
	}
)

// SearchHandler search needs for the term
func SearchHandler(sR SearchNeedRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryValues := r.URL.Query()

		var orgID int64
		var categoriesID []int
		var orderBy []string
		var err error

		if len(queryValues.Get("text")) < 1 || len(queryValues.Get("page")) < 1 {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Parametros inválidos"))
			return
		}

		if len(queryValues.Get("org")) > 0 {
			orgID, err = strconv.ParseInt(queryValues.Get("org"), 10, 64)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("org")))
			}
		}

		if len(queryValues.Get("categories")) > 0 {
			IDsplited := strings.Split(queryValues.Get("categories"), ",")
			categoriesID = make([]int, len(IDsplited))
			for i, v := range IDsplited {
				categoriesID[i], err = strconv.Atoi(v)
				if err != nil {
					HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", v))
				}
			}
		}

		page, err := strconv.ParseInt(queryValues.Get("page"), 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("page")))
		}

		if len(queryValues.Get("order_by")) > 0 {
			orderBy = strings.Split(queryValues.Get("order_by"), ",")
		}

		needs, err := sR.Search(queryValues.Get("text"), categoriesID, orgID, orderBy, page)

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
