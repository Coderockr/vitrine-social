package handlers

import (
	"database/sql"
	"encoding/json"
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

// NeedHandler handles requests about Needs
type NeedHandler struct {
	repo  needRepository
	oRepo OrganizationRepository
}

// NewNeedHandler creates a new NeedHandler
func NewNeedHandler(repo needRepository, oRepo needOrganizationRepository) NeedHandler {
	return NeedHandler{
		repo:  repo,
		oRepo: oRepo,
	}
}

// NeedGet retorna uma necessidade pelo ID
func (h NeedHandler) NeedGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHttpError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		n, err := h.repo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHttpError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
			return
		case err != nil:
			HandleHttpError(w, http.StatusForbidden, err)
			return
		}

		o, err := h.oRepo.Get(n.OrganizationID)
		if err != nil {
			HandleHttpError(w, http.StatusForbidden, err)
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

		if err := json.NewEncoder(w).Encode(nJSON); err != nil {
			HandleHttpError(w, http.StatusInternalServerError, err)
			return
		}
	})
}
