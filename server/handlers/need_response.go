package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Coderockr/vitrine-social/server/mail"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gorilla/mux"
)

type needResponseRepository interface {
	CreateResponse(*model.NeedResponse) (int64, error)
}

// NeedResponse responde uma necessidade pelo ID
func NeedResponse(needRepo NeedRepository, needResponseRepo needResponseRepository, mailer mail.Mailer) func(w http.ResponseWriter, r *http.Request) {
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

		var bodyVars struct {
			Name    string
			Email   string
			Phone   string
			Address string
			Message string
		}

		err = requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		if bodyVars.Email == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'email' é obrigatório! "))
			return
		}

		if bodyVars.Name == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'name' é obrigatório! "))
			return
		}

		if bodyVars.Phone == "" {
			HandleHTTPError(w, http.StatusBadRequest, fmt.Errorf(" O campo 'phone' é obrigatório! "))
			return
		}

		newID, err := needResponseRepo.CreateResponse(&model.NeedResponse{
			Email:   bodyVars.Email,
			Name:    bodyVars.Name,
			Phone:   bodyVars.Phone,
			Address: bodyVars.Address,
			Message: bodyVars.Message,
			NeedID:  n.ID,
		})

		if err != nil {
			HandleHTTPError(w, http.StatusForbidden, err)
			return
		}

		emailParams := mail.EmailParams{
			To:       n.Organization.User.Email,
			Subject:  "Vitrine Social - Resposta de Solicitação",
			Template: mail.NeedResponseTemplate,
			Variables: map[string]string{
				"need":    n.Title,
				"name":    bodyVars.Name,
				"email":   bodyVars.Email,
				"phone":   bodyVars.Phone,
				"address": bodyVars.Address,
				"message": bodyVars.Message,
			},
		}

		if err := mailer.SendEmail(emailParams); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao enviar o email"))
			return
		}

		HandleHTTPSuccess(w, map[string]int64{"id": newID})
	}
}
