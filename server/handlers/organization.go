package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type (
	// OrganizationRepository represet operations for organization repository.
	OrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
		Update(o model.Organization) (model.Organization, error)
		DeleteImage(imageID int64, organizationID int64) error
	}
)

// GetOrganizationHandler will retrive the data from a organization
func GetOrganizationHandler(getOrg func(int64) (*model.Organization, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		o, err := getOrg(id)

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
			addressJSON: addressJSON{
				Street:     o.Street,
				Number:     o.Number,
				Complement: o.Complement,
				Suburb:     o.Suburb,
				City:       o.City,
				State:      o.State,
				Zipcode:    o.Zipcode,
			},
			Phone:  o.Phone,
			Resume: o.Resume,
			Video:  o.Video,
			Email:  o.Email,
			Images: orgImagesToImageJSON(o.Images),
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
				Status:           string(n.Status),
			})
		}
		HandleHTTPSuccess(w, oJSON)
	}
}

// UpdateOrganizationHandler will update the data of an organization
func UpdateOrganizationHandler(repo OrganizationRepository) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var organization model.Organization

		err := requestToJSONObject(req, &organization)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		vars := mux.Vars(req)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		userID := GetUserID(req)
		if id == 0 || id != userID {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("você não possui permissão para atualizar essa orginização %d", id))
			return
		}

		_, err = repo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Organização com ID: %d", id))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		organization.ID = id

		_, err = repo.Update(organization)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar dados da Organização: %s", err))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}

// DeleteOrganizationImageHandler will delete the image
func DeleteOrganizationImageHandler(repo OrganizationRepository) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		HandleHTTPSuccessNoContent(w)
	}
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
