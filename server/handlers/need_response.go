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
	CreateResponse(*model.NeedResponse) (sql.Result, error)
}

// NeedResponse responde uma necessidade pelo ID
func NeedResponse(needRepo needRepository, needResponseRepo needResponseRepository) interface{} {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf("Não foi possível entender o número: %s", vars["id"]))
			return
		}
		email := vars["email"]
		name := vars["name"]
		phone := vars["phone"]
		address := vars["address"]
		message := vars["message"]

		n, err := needRepo.Get(id)
		switch {
		case err == sql.ErrNoRows:
			HandleHTTPError(w, http.StatusNotFound, fmt.Errorf("Não foi encontrada Necessidade com ID: %d", id))
			return
		case err != nil:
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}
		now := time.Now()
		newID, err := needResponseRepo.CreateResponse(&model.NeedResponse{
			Date:    &now,
			Email:   email,
			Name:    name,
			Phone:   phone,
			Address: address,
			Message: message,
			NeedID:  n.ID,
		})
		if err != nil {
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}
		HandleHTTPSuccess(w, newID)

	}
}
