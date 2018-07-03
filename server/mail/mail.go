package mail

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Mailer is a implementation to send emails
type Mailer struct {
	Dial gomail.SendCloser
}

// EmailParams struct with emails infos
type EmailParams struct {
	From    string
	To      string
	Subject string
	Body    string
}

// Connect - Create and return a dialer
func Connect() (Mailer, error) {
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	d := gomail.NewPlainDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	s, err := d.Dial()

	return Mailer{
		Dial: s,
	}, err
}

// SendEmail - Send email
func (d Mailer) SendEmail(params EmailParams) error {
	m := gomail.NewMessage()
	m.SetHeader("From", params.From)
	m.SetHeader("To", params.To)
	m.SetHeader("Subject", params.Subject)
	m.SetBody("text/html", params.Body)

	return gomail.Send(d.Dial, m)
}
