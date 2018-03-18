package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type needRepository interface {
	Get(id int64) (*model.Need, error)
}

type needOrganizationRepository interface {
	Get(id int64) (*model.Organization, error)
}

// GetNeedHandler retorna uma necessidade pelo ID
func GetNeedHandler(repo needRepository, oRepo needOrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		n, err := repo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		o, err := oRepo.Get(n.OrganizationID)
		if err != nil {
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		var dueDate *jsonTime
		if n.DueDate != nil {
			dueDate = &jsonTime{*n.DueDate}
		}
		nJSON := needJSON{
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
				ID:   o.ID,
				Name: o.Name,
				Logo: o.Logo,
				Slug: o.Slug,
			},
			Images: needImagesToImageJSON(n.Images),
			Status: n.Status,
		}

		HandleHTTPSuccess(w, nJSON)
	}
}
