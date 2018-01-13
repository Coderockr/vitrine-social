package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

// OrganizationRepository has the commands needed for this route
type OrganizationRepository interface {
	Get(id int64) (*model.Organization, error)
}

// OrganizationHandler handles requests about organizations
type OrganizationHandler struct {
	repo OrganizationRepository
}

// NewOrganizationHandler creates a new OrganizationHandler
func NewOrganizationHandler(repo OrganizationRepository) *OrganizationHandler {
	return &OrganizationHandler{
		repo: repo,
	}
}

// Get will retrive the data from a organization
func (oR *OrganizationHandler) Get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
		return
	}

	o, err := oR.repo.Get(id)

	switch {
	case err == sql.ErrNoRows:
		HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Organização com ID: %d", id))
		return
	case err != nil:
		HandleHTTPError(w, http.StatusForbidden, err)
		return
	}

	oJSON := &organizationJSON{
		baseOrganizationJSON: baseOrganizationJSON{
			ID:   o.ID,
			Name: o.Name,
			Logo: o.Logo,
			Slug: o.Slug,
		},
		Address: o.Address,
		Phone:   o.Phone,
		Resume:  o.Resume,
		Video:   o.Video,
		Email:   o.Email,
		Images:  orgImagesToImageJSON(o.Images),
	}

	oJSON.Needs = make([]needJSON, 0, len(o.Needs))
	catMap := make(map[int64]categoryJSON)

	for _, n := range o.Needs {
		if _, ok := catMap[n.CategoryID]; ok == false {
			catMap[n.CategoryID] = categoryJSON{
				ID:   n.Category.ID,
				Name: n.Category.Name,
				Icon: n.Category.Icon,
			}
		}

		var dueDate *jsonTime
		if n.DueDate != nil {
			dueDate = &jsonTime{*n.DueDate}
		}
		oJSON.Needs = append(oJSON.Needs, needJSON{
			ID:               n.ID,
			Title:            n.Title,
			Description:      n.Description,
			RequiredQuantity: n.RequiredQuantity,
			ReachedQuantity:  n.ReachedQuantity,
			Unity:            n.Unity,
			DueDate:          dueDate,
			Category:         catMap[n.CategoryID],
			Organization:     oJSON.baseOrganizationJSON,
			Images:           needImagesToImageJSON(n.Images),
			Status:           n.Status,
		})
	}
	HandleHTTPSuccess(w, oJSON)
}

func needImagesToImageJSON(images []model.NeedImage) []imageJSON {
	imagesJSON := make([]imageJSON, 0, len(images))
	for _, i := range images {
		imagesJSON = append(imagesJSON, imageJSON{
			ID:   i.ID,
			Name: i.Name,
			URL:  i.URL,
		})
	}
	return imagesJSON
}

func orgImagesToImageJSON(images []model.OrganizationImage) []imageJSON {
	imagesJSON := make([]imageJSON, 0, len(images))
	for _, i := range images {
		imagesJSON = append(imagesJSON, imageJSON{
			ID:   i.ID,
			Name: i.Name,
			URL:  i.URL,
		})
	}
	return imagesJSON
}
