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

	from := mail.NewEmail("Vitrine Social", os.Getenv("MAIL_FROM"))
	to := mail.NewEmail("", params.To)
	message := mail.NewV3MailInit(from, params.Subject, to)

	if params.Template != "" {
		switch params.Template {
		case ForgotPasswordTemplate:
			message.SetTemplateID(os.Getenv("SENDGRID_TEMPLATE_FORGOT_PASSWORD"))
		case NeedResponseTemplate:
			message.SetTemplateID(os.Getenv("SENDGRID_TEMPLATE_NEED_RESPONSE"))
		}
	}

	if len(params.Variables) > 0 {
		for i := range params.Variables {
			message.Personalizations[0].SetSubstitution("{{"+i+"}}", params.Variables[i])
		}
	}

	_, err = mailer.Client.Send(message)

	return err
}
