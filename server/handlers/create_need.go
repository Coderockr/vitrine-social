package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

		cat, err := getCat(catID)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada categoria com ID: %d", catID))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		org, err := getOrg(orgID)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Organização com ID: %d", catID))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		var title, description string
		var requiredQuantity, reachedQuantity int
		var dueDate *time.Time
		var ok bool

		if title, ok = bodyVars["title"]; !ok || len(strings.TrimSpace(title)) == 0 {
			HandleHTTPError(w, http.StatusForbidden, fmt.Errorf("Deve ser informado um título para a Necessidade"))
			return
		}

		if description, ok = bodyVars["description"]; !ok || len(strings.TrimSpace(description)) == 0 {
			HandleHTTPError(w, http.StatusForbidden, fmt.Errorf("Deve ser informado um descrição para a Necessidade"))
			return
		}

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
			OrganizationID:   org.ID,
			CategoryID:       cat.ID,
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
