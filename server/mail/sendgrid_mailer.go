package mail

import (
	"fmt"
	"log"
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

	from := mail.NewEmail("Vitrine Social", os.Getenv("MAIL_FROM"))
	to := mail.NewEmail("", params.To)
	message := mail.NewSingleEmail(from, params.Subject, to, params.Body, params.Body)

	if params.TemplateID != "" {
		message.SetTemplateID(params.TemplateID)
	}

	if len(params.Variables) > 0 {
		personalization := mail.NewPersonalization()
		personalization.AddTos(to)
		for i := range params.Variables {
			personalization.SetSubstitution(i, params.Variables[i])
		}
		message = message.AddPersonalizations(personalization)
	}

	response, err := mailer.Client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return err
}
