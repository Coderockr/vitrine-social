package handlers

import (
	"errors"
	"net/http"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
)

// NewsletterParams newsletter contact params
type NewsletterParams struct {
	Name  string
	Email string
	Phone string
}

// NewsletterHandler add new email to newsletter list
func NewsletterHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyVars NewsletterParams
		err := requestToJSONObject(r, &bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		err = SaveNewsletter(bodyVars)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao salvar novo contato"))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}

// SaveNewsletter send request to save new contact to sendgrid
func SaveNewsletter(params NewsletterParams) error {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/contactdb/recipients", "")
	request.Method = "POST"
	request.Body = []byte(`[
		{
			"first_name": "` + params.Name + `",
			"email": "` + params.Email + `",
			"phone": "` + params.Phone + `"
		}
	]`)
	_, err := sendgrid.API(request)
	return err
}
