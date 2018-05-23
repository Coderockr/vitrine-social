package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type (
	// SubscriptionRepository represet operations for organization repository.
	SubscriptionRepository interface {
		Create(model.Subscription) (model.Subscription, error)
	}

	subscriptionOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
	}
)

// CreateSubscriptionHandler create a new subscription
func CreateSubscriptionHandler(repo SubscriptionRepository, orgRepo subscriptionOrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlVars := mux.Vars(r)
		id, err := strconv.ParseInt(urlVars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", urlVars["id"]))
			return
		}

		o, err := orgRepo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Organização com ID: %d", id))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		var bodyVars map[string]string
		err = requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		if bodyVars["email"] == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'email' é obrigatório! "))
			return
		}

		if bodyVars["name"] == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'name' é obrigatório! "))
			return
		}

		if bodyVars["phone"] == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'phone' é obrigatório! "))
			return
		}

		now := time.Now()
		s, err := repo.Create(model.Subscription{
			OrganizationID: o.ID,
			Email:          bodyVars["email"],
			Name:           bodyVars["name"],
			Phone:          bodyVars["phone"],
			Date:           &now,
		})

		if err != nil {
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": s.ID})
	}
}
