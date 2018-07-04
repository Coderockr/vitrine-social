package mail

import (
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridMailer is a implementation of SendGrid API
type SendGridMailer struct {
	Client       *sendgrid.Client
	mailSettings *mail.MailSettings
}

// SendgridConnect - Create and return a mailer
func SendgridConnect() (Mailer, error) {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	var mailSettings *mail.MailSettings

	if os.Getenv("MAIL_SANDBOX") == "true" {
		mailSettings = mail.NewMailSettings()
		mailSettings.SetSandboxMode(mail.NewSetting(true))
	}

	return SendGridMailer{Client: client, mailSettings: mailSettings}, nil
}

// SendEmail - Send email with SendGridMailer
func (mailer SendGridMailer) SendEmail(params EmailParams) error {
	var err error

	from := mail.NewEmail("", params.From)
	to := mail.NewEmail("", params.To)
	message := mail.NewSingleEmail(from, params.Subject, to, params.Body, params.Body)

	if params.TemplateID != "" {
		message.SetTemplateID(params.TemplateID)
	}

	_, err = mailer.Client.Send(message)

	return err
}
