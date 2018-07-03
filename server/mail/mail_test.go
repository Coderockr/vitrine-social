package mail_test

import (
	"log"
	"testing"

	"github.com/Coderockr/vitrine-social/server/mail"
	"github.com/joho/godotenv"
)

type (
	mailMock struct {
		ConnectFN   func() (mail.Mailer, error)
		SendEmailFN func(mail.EmailParams) error
	}
)

func TestConnect(t *testing.T) {
	err := godotenv.Load("../config/test.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// _, err = mail.Connect()
	// if err != nil {
	// 	t.Errorf("expected %v result %s", nil, err)
	// }
}

func (m *mailMock) Connect() (mail.Mailer, error) {
	return m.ConnectFN()
}

func (m *mailMock) SendEmail(params mail.EmailParams) error {
	return m.SendEmailFN(params)
}
