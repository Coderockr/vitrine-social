package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/Coderockr/vitrine-social/server/mail"
	"github.com/Coderockr/vitrine-social/server/newsletter"
)

// ContactHandler send an email to coderockr
func ContactHandler(mailer mail.Mailer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars struct {
			Name       string
			Email      string
			Phone      string
			Reason     string
			Message    string
			Newsletter bool
		}
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		emailParams := mail.EmailParams{
			To:       os.Getenv("MAIL_CONTACT"),
			Subject:  "Vitrine Social - Contato",
			Template: mail.ContactTemplate,
			Variables: map[string]string{
				"name":    bodyVars.Name,
				"email":   bodyVars.Email,
				"phone":   bodyVars.Phone,
				"reason":  bodyVars.Reason,
				"message": bodyVars.Message,
			},
		}

		if err := mailer.SendEmail(emailParams); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao enviar o email"))
			return
		}

		if bodyVars.Newsletter {
			newsletterParams := newsletter.Params{
				Name:  bodyVars.Name,
				Email: bodyVars.Email,
				Phone: bodyVars.Phone,
			}
			err = newsletter.SaveNewsletter(newsletterParams)
			if err != nil {
				HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao salvar novo contato na newsletter"))
				return
			}
		}

		HandleHTTPSuccessNoContent(w)
	}
}
