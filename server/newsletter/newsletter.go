package newsletter

import (
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
)

// NewsletterParams newsletter contact params
type NewsletterParams struct {
	Name  string
	Email string
	Phone string
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
