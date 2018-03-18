package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
)

// CreateNeedHandler creates a new need based on the payload
func CreateNeedHandler(
	create func(model.Need) (model.Need, error),
	getOrg func(id int64) (*model.Organization, error),
	getCat func(id int64) (model.Category, error),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars map[string]string
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		userID := GetUserID(r)
		orgID, err := strconv.ParseInt(bodyVars["organization"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("não foi possível reconhecer o ID da organização"))
			return
		}

		if orgID == 0 || userID != orgID {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("você não possui permissão para criar uma necessidade para a orginização %d", orgID))
			return
		}

		catID, err := strconv.ParseInt(bodyVars["category"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("não foi possível reconhecer o ID da organização"))
			return
		}

		var requiredQuantity, reachedQuantity int
		var dueDate *time.Time

		title := bodyVars["title"]
		description := bodyVars["description"]

		if str, ok := bodyVars["requiredQuantity"]; ok {
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível reconhecer quantidade requerida informada"))
				return
			}
			requiredQuantity = int(i)
		}

		if str, ok := bodyVars["reachedQuantity"]; ok {
			i, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível reconhecer quantidade alcançada informada"))
				return
			}
			reachedQuantity = int(i)
		}

		if str, ok := bodyVars["dueDate"]; ok {
			t, err := time.Parse("2006-01-02", str)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível reconhecer a data informada"))
				return
			}
			dueDate = &t
		}

		n, err := create(model.Need{
			OrganizationID:   orgID,
			CategoryID:       catID,
			Title:            title,
			Description:      description,
			RequiredQuantity: requiredQuantity,
			ReachedQuantity:  reachedQuantity,
			DueDate:          dueDate,
		})

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": n.ID})
	}
}
