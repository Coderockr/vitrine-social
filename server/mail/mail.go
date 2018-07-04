package mail

import (
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

// Mailer is a implementation to send emails
type Mailer struct {
	Dial   *gomail.SendCloser
	Client *sendgrid.Client
}

// EmailParams struct with emails infos
type EmailParams struct {
	From       string
	To         string
	Subject    string
	Body       string
	TemplateID string
}

// Connect - Create and return a dialer
func Connect() (Mailer, error) {
	method := os.Getenv("MAIL_METHOD")
	var mailer Mailer
	var err error

	switch method {
	case "sendgrid":
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		mailer = Mailer{
			Dial:   nil,
			Client: client,
		}
		err = nil

	case "smtp":
		var dial gomail.SendCloser
		mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
		dialer := gomail.NewPlainDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
		dial, err = dialer.Dial()

		mailer = Mailer{
			Dial:   &dial,
			Client: nil,
		}
	}

	return mailer, err
}

// SendEmail - Send email
func (mailer Mailer) SendEmail(params EmailParams) error {
	var err error
	method := os.Getenv("MAIL_METHOD")

	switch method {
	case "sendgrid":
		from := mail.NewEmail("", params.From)
		to := mail.NewEmail("", params.To)
		message := mail.NewSingleEmail(from, params.Subject, to, params.Body, params.Body)

		if os.Getenv("MAIL_SANDBOX") == "true" {
			mailSettings := mail.NewMailSettings()
			mailSettings.SetSandboxMode(mail.NewSetting(true))
			message.SetMailSettings(mailSettings)
		}

		if params.TemplateID != "" {
			message.SetTemplateID(params.TemplateID)
		}

		_, err = mailer.Client.Send(message)

	case "smtp":
		message := gomail.NewMessage()
		message.SetHeader("From", params.From)
		message.SetHeader("To", params.To)
		message.SetHeader("Subject", params.Subject)
		message.SetBody("text/html", params.Body)

		err = gomail.Send(*mailer.Dial, message)
	}

	return err
}
