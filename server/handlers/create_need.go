package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
)

// CreateNeedHandler creates a new need based on the payload
func CreateNeedHandler(create func(model.Need) (model.Need, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars struct {
			Organization     int64
			Category         int64
			Title            string
			Description      string
			RequiredQuantity int
			ReachedQuantity  int
			Unit             string
			DueDate          *jsonTime
		}
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		userID := GetUserID(r)
		if bodyVars.Organization == 0 || userID != bodyVars.Organization {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("você não possui permissão para criar uma necessidade para a orginização %d", bodyVars.Organization))
			return
		}

		var dueDate *time.Time
		if bodyVars.DueDate != nil {
			dueDate = &bodyVars.DueDate.Time
		}

		n, err := create(model.Need{
			OrganizationID:   bodyVars.Organization,
			CategoryID:       bodyVars.Category,
			Title:            bodyVars.Title,
			Description:      bodyVars.Description,
			RequiredQuantity: bodyVars.RequiredQuantity,
			ReachedQuantity:  bodyVars.ReachedQuantity,
			Unit:             bodyVars.Unit,
			DueDate:          dueDate,
		})

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": n.ID})
	}
}
