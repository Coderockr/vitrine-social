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
				ID:    o.ID,
				Name:  o.Name,
				Logo:  o.Logo,
				Slug:  o.Slug,
				Phone: o.Phone,
			},
			Address: addressJSON{
				Street:       o.Address.Street,
				Number:       o.Address.Number,
				Complement:   o.Address.Complement,
				Neighborhood: o.Address.Neighborhood,
				City:         o.Address.City,
				State:        o.Address.State,
				Zipcode:      o.Address.Zipcode,
			},
			About:  o.About,
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
					Slug: n.Category.Slug,
				}
			}

			var dueDate *jsonTime
			var createdAt *jsonTime
			var updatedAt *jsonTime

			if n.DueDate != nil {
				dueDate = &jsonTime{*n.DueDate}
			}
			if n.CreatedAt != nil {
				createdAt = &jsonTime{*n.CreatedAt}
			}
			if n.UpdatedAt != nil {
				updatedAt = &jsonTime{*n.UpdatedAt}
			}
			oJSON.Needs = append(oJSON.Needs, needJSON{
				ID:               n.ID,
				Title:            n.Title,
				Description:      n.Description,
				RequiredQuantity: n.RequiredQuantity,
				ReachedQuantity:  n.ReachedQuantity,
				Unit:             n.Unit,
				DueDate:          dueDate,
				CreatedAt:        createdAt,
				UpdatedAt:        updatedAt,
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
		vars := mux.Vars(req)
		organizationID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		imageID, err := strconv.ParseInt(vars["image_id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["image_id"]))
			return
		}

		userID := GetUserID(req)
		if organizationID == 0 || organizationID != userID {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("você não possui permissão para remover a Imagem %d", imageID))
			return
		}

		err = repo.DeleteImage(imageID, organizationID)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Imagem com ID: %d", imageID))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao deletar Imagem: %s", err))
			return
		}

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
