package handlers

import (
	"errors"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/newsletter"
)

// NewsletterHandler add new email to newsletter list
func NewsletterHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars newsletter.Params
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		err = newsletter.SaveNewsletter(bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao salvar novo contato"))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}
