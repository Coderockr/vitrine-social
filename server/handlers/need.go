package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type (
	// NeedRepository represet operations for need repository.
	NeedRepository interface {
		Get(id int64) (*model.Need, error)
		Update(model.Need) (model.Need, error)
	}

	needOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
	}
)

// GetNeedHandler retorna uma necessidade pelo ID
func GetNeedHandler(repo NeedRepository, oRepo needOrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
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
			Status: string(*n.Status),
		}

		HandleHTTPSuccess(w, nJSON)
	}
}

// UpdateNeedHandler get the need, update and save on database
func UpdateNeedHandler(repo NeedRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars struct {
			Category         int64
			Title            string
			Description      string
			RequiredQuantity int
			ReachedQuantity  int
			DueDate          *jsonTime
			Unity            string
		}
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		need, err := repo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		userID := GetUserID(r)
		if need.OrganizationID != userID {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("você não possui permissão para atualizar essa necessidade %d", need.OrganizationID))
			return
		}

		var dueDate *time.Time
		if bodyVars.DueDate != nil {
			dueDate = &bodyVars.DueDate.Time
		}

		need.CategoryID = bodyVars.Category
		need.Title = bodyVars.Title
		need.Description = bodyVars.Description
		need.RequiredQuantity = bodyVars.RequiredQuantity
		need.ReachedQuantity = bodyVars.ReachedQuantity
		need.DueDate = dueDate
		need.Unity = bodyVars.Unity

		_, err = repo.Update(*need)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar dados da necessidade: %s", err))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}
