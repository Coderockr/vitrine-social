package mail

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/gomail.v2"
)

// Mailer is a implementation to send emails
type Mailer interface {
	SendEmail(EmailParams) error
}

// SendGridMailer is a implementation of SendGrid API
type SendGridMailer struct {
	Client       *sendgrid.Client
	mailSettings *mail.MailSettings
}

// SMTPMailer is a implementation of SMTP
type SMTPMailer struct {
	Dial *gomail.SendCloser
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

	switch method {
	case "sendgrid":
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		var mailSettings *mail.MailSettings

		if os.Getenv("MAIL_SANDBOX") == "true" {
			mailSettings = mail.NewMailSettings()
			mailSettings.SetSandboxMode(mail.NewSetting(true))
		}

		return SendGridMailer{Client: client, mailSettings: mailSettings}, nil

	case "smtp":
		var dial gomail.SendCloser
		mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
		dialer := gomail.NewPlainDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
		dial, err := dialer.Dial()

		return SMTPMailer{Dial: &dial}, err

	default:
		return nil, fmt.Errorf("mail method %s is not known", method)
	}
}

// SendEmail - Send email
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

// SendEmail - Send email
func (mailer SMTPMailer) SendEmail(params EmailParams) error {
	message := gomail.NewMessage()
	message.SetHeader("From", params.From)
	message.SetHeader("To", params.To)
	message.SetHeader("Subject", params.Subject)
	message.SetBody("text/html", params.Body)

	return gomail.Send(*mailer.Dial, message)
}
