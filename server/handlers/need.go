package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gobuffalo/uuid"
	"github.com/gorilla/mux"
	"github.com/graymeta/stow"
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
		Put(name string, r io.Reader, size int64, metadata map[string]interface{}) (stow.Item, error)
		RemoveItem(id string) error
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
			Status: string(n.Status),
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

// UploadNeedImagesHandler upload file to storage and save new image
func UploadNeedImagesHandler(repo NeedRepository, container needStorageContainer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}

		file, handler, err := r.FormFile("images")
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível ler o arquivo"))
			return
		}
		defer file.Close()

		fileName := strings.Split(handler.Filename, ".")

		uuid := uuid.Must(uuid.NewV4())
		path := "need-" + vars["id"] + "/" + uuid.String() + "." + fileName[1]
		item, err := container.Put(path, file, handler.Size, nil)
		if err != nil {
			log.Fatalf("Erro ao salvar arquivo: %v\n", err)
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar arquivo"))
			return
		}

		image := model.NeedImage{
			Image: model.Image{
				Name: fileName[0],
				URL:  item.ID(),
			},
			NeedID: id,
		}

		image, err = repo.CreateImage(image)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Erro ao salvar imagem: %s", err))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}
