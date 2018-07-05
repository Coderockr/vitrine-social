package handlers

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

type (
	// NeedRepository represet operations for need repository.
	NeedRepository interface {
		Get(id int64) (*model.Need, error)
		Update(model.Need) (model.Need, error)
		CreateImage(i model.NeedImage) (model.NeedImage, error)
	}

	needOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
	}

	needStorageContainer interface {
		CreateNeedImage(*model.Token, int64, *multipart.FileHeader) (*model.NeedImage, error)
		DeleteNeedImage(t *model.Token, nID, iID int64) error
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
			Unit:             n.Unit,
			DueDate:          dueDate,
			Category: categoryJSON{
				ID:   n.Category.ID,
				Name: n.Category.Name,
				Slug: n.Category.Slug,
			},
			Organization: baseOrganizationJSON{
				ID:   o.ID,
				Name: o.Name,
				Logo: o.Logo,
				Slug: o.Slug,
			},
			Images:    needImagesToImageJSON(n.Images),
			Status:    string(n.Status),
			CreatedAt: n.CreatedAt,
			UpdatedAt: n.UpdatedAt,
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
			Unit             string
			Status           string
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

		err = need.Status.Scan(bodyVars.Status)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("%s não é um status válido: %s", bodyVars.Status, err))
			return
		}

		need.CategoryID = bodyVars.Category
		need.Title = bodyVars.Title
		need.Description = bodyVars.Description
		need.RequiredQuantity = bodyVars.RequiredQuantity
		need.ReachedQuantity = bodyVars.ReachedQuantity
		need.DueDate = dueDate
		need.Unit = bodyVars.Unit

		_, err = repo.Update(*need)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar dados da necessidade: %s", err))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}

// UploadNeedImagesHandler upload file to storage and save new image
func UploadNeedImagesHandler(container needStorageContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		if err := r.ParseMultipartForm(defaultMaxMemory); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		files := r.MultipartForm.File["images"]
		if len(files) == 0 {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível ler o arquivo"))
			return
		}

		t := GetModelToken(r)
		i, err := container.CreateNeedImage(t, id, files[0])
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar imagem: %s", err))
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": i.ID})
	}
}

// DeleteNeedImagesHandler will delete the image
func DeleteNeedImagesHandler(storage needStorageContainer) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		needID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		imageID, err := strconv.ParseInt(vars["image_id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["image_id"]))
			return
		}

		t := GetModelToken(req)
		if err = storage.DeleteNeedImage(t, needID, imageID); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}
