package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// SearchNeedRepository represet operations for need repository.
	SearchNeedRepository interface {
		Search(text string, categoriesID []int, organizationsID int64, status string, orderBy string, order string, page int) ([]model.SearchNeed, int, error)
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

		page, err := strconv.Atoi(queryValues.Get("page"))
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", queryValues.Get("page")))
		}

		text := queryValues.Get("text")
		status := queryValues.Get("status")
		orderBy := queryValues.Get("orderBy")
		order := queryValues.Get("order")

		needs, count, err := sR.Search(text, categoriesID, orgID, status, orderBy, order, page)

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, convertDBToNeed(count, page, needs))
	}
}

func convertDBToNeed(count int, currentPage int, searchNeed []model.SearchNeed) searchResultJSON {
	need := searchResultJSON{
		Pagination: paginationJSON{
			TotalResults: count,
			TotalPages:   int(math.Ceil(float64(count) / repo.ResultsPerPage)),
			CurrentPage:  currentPage,
		},
		Results: make([]needJSON, len(searchNeed)),
	}

	var dueDate *jsonTime
	for i, s := range searchNeed {
		if s.DueDate != nil {
			dueDate = &jsonTime{*s.DueDate}
		}

		need.Results[i] = needJSON{
			ID:               s.ID,
			Title:            s.Title,
			Description:      s.Description,
			RequiredQuantity: s.RequiredQuantity,
			ReachedQuantity:  s.ReachedQuantity,
			Unit:             s.Unit,
			DueDate:          dueDate,
			CreatedAt:        s.CreatedAt,
			UpdatedAt:        s.UpdatedAt,
			Images:           needImagesToImageJSON(s.Images),
			Category: categoryJSON{
				ID:   s.CategoryID,
				Name: s.CategoryName,
				Slug: s.CategorySlug,
			},
			Organization: baseOrganizationJSON{
				ID:   s.OrganizationID,
				Name: s.OrganizationName,
				Logo: imageJSON{
					ID:   s.OrganizationLogo.ID,
					Name: s.OrganizationLogo.Name,
					URL:  s.OrganizationLogo.URL,
				},
				Slug:  s.OrganizationSlug,
				Phone: s.OrganizationPhone,
			},
		}
	}

	return need
}
