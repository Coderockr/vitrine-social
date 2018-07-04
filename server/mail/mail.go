package mail

import (
	"fmt"
	"os"
)

// Mailer is a implementation to send emails
type Mailer interface {
	SendEmail(EmailParams) error
}

// EmailParams struct with emails infos
type EmailParams struct {
	To         string
	Subject    string
	Body       string
	TemplateID string
	Variables  map[string]string
}

// Connect - Create and return a dialer
func Connect() (Mailer, error) {
	method := os.Getenv("MAIL_METHOD")

	switch method {
	case "sendgrid":
		return SendgridConnect()

	case "smtp":
		return SMTPConnect()

	default:
		return nil, fmt.Errorf("mail method %s is not known", method)
	}
}
