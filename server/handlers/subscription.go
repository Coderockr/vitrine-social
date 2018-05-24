package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type (
	// SubscriptionRepository represet operations for subscription repository.
	SubscriptionRepository interface {
		Create(model.Subscription) (model.Subscription, error)
	}
)

// CreateSubscriptionHandler create a new subscription
func CreateSubscriptionHandler(repo SubscriptionRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlVars := mux.Vars(r)
		id, err := strconv.ParseInt(urlVars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", urlVars["id"]))
			return
		}

		var bodyVars map[string]string
		err = requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		now := time.Now()
		s, err := repo.Create(model.Subscription{
			OrganizationID: id,
			Email:          bodyVars["email"],
			Name:           bodyVars["name"],
			Phone:          bodyVars["phone"],
			Date:           &now,
		})

		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": s.ID})
	}
}
