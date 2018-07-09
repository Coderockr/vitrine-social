package mail

import (
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

// SMTPMailer is a implementation of SMTP
type SMTPMailer struct {
	Dial *gomail.SendCloser
}

// SMTPConnect - Create and return a dialer
func SMTPConnect() (Mailer, error) {
	var dial gomail.SendCloser
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	dialer := gomail.NewPlainDialer(os.Getenv("MAIL_HOST"), mailPort, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))
	dial, err := dialer.Dial()

	return SMTPMailer{Dial: &dial}, err
}

// SendEmail - Send email with SMTPMailer
func (mailer SMTPMailer) SendEmail(params EmailParams) error {
	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("MAIL_FROM"))
	message.SetHeader("To", params.To)
	message.SetHeader("Subject", params.Subject)
	message.SetBody("text/html", "")

	return gomail.Send(*mailer.Dial, message)
}
