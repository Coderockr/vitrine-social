package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/db/repo"
	"github.com/Coderockr/vitrine-social/server/index"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

// NeedHandler handles requests about Needs
type NeedHandler struct {
	repo         *repo.NeedRepository
	oRepo        *repo.OrganizationRepository
	indexService index.Service
}

// NewNeedHandler creates a new NeedHandler
func NewNeedHandler(repo *repo.NeedRepository, oRepo *repo.OrganizationRepository, indexService index.Service) NeedHandler {
	return NeedHandler{
		repo:         repo,
		oRepo:        oRepo,
		indexService: indexService,
	}
}

// NeedGet retorna uma necessidade pelo ID
func (h NeedHandler) NeedGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
		return
	}

	n, err := h.repo.Get(id)
	switch {
	case err == sql.ErrNoRows:
		HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
		return
	case err != nil:
		HandleHTTPError(w, http.StatusForbidden, err)
		return
	}

	o, err := h.oRepo.Get(n.OrganizationID)
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

// NeedPost adiciona uma need
func (h NeedHandler) NeedPost(w http.ResponseWriter, r *http.Request) {
	var nj needJSON
	err := json.NewDecoder(r.Body).Decode(&nj)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, err)
		return
	}
	o, err := h.oRepo.Get(nj.Organization.ID)
	if err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}
	//@todo validate data
	n := &model.Need{
		Title:            nj.Title,
		Description:      nj.Description,
		RequiredQuantity: nj.RequiredQuantity,
		ReachedQuantity:  nj.ReachedQuantity,
		Unity:            nj.Unity,
		DueDate:          &nj.DueDate.Time,
		CategoryID:       nj.Category.ID,
		OrganizationID:   nj.Organization.ID,
	}

	id, err := h.repo.Insert(n)
	if err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}
	key := strconv.FormatInt(id, 10)
	toIndex := model.NeedToIndex{
		Key:   key,
		ID:    id,
		Title: n.Title,
		// CategoryName:       n.Category.Name,
		OrganizationName:   o.Name,
		OrganizationResume: o.Resume,
		OrganizationSlug:   o.Slug,
	}
	d := index.Data{
		Key:  key,
		ID:   id,
		Data: toIndex,
	}
	err = h.indexService.Index(key, d)
	if err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}
	type data struct {
		ID int64 `json:"id"`
	}
	HandleHTTPSuccess(w, data{ID: id}, http.StatusCreated)
}
