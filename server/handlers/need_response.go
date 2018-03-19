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

type needResponseRepository interface {
	CreateResponse(*model.NeedResponse) (int64, error)
}

// NeedResponse responde uma necessidade pelo ID
func NeedResponse(needRepo needRepository, needResponseRepo needResponseRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlVars := mux.Vars(r)
		id, err := strconv.ParseInt(urlVars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", urlVars["id"]))
			return
		}

		n, err := needRepo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
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
		newID, err := needResponseRepo.CreateResponse(&model.NeedResponse{
			Date:    &now,
			Email:   bodyVars["email"],
			Name:    bodyVars["name"],
			Phone:   bodyVars["phone"],
			Address: bodyVars["address"],
			Message: bodyVars["message"],
			NeedID:  n.ID,
		})
		if err != nil {
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": newID})

	}

}
