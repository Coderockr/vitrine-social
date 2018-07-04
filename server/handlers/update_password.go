package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Coderockr/vitrine-social/server/mail"
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/security"
)

type (
	// UpdatePasswordOrganizationRepository represet operations for organization repository.
	UpdatePasswordOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
		GetByEmail(email string) (*model.Organization, error)
		ResetPasswordTo(o *model.Organization, password string) error
	}
)

// UpdatePasswordHandler update user password
func UpdatePasswordHandler(repo UpdatePasswordOrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var handlerForm map[string]string

		err := requestToJSONObject(r, &handlerForm)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		userID := GetUserID(r)
		organization, err := repo.Get(userID)

		err = security.CompareHashAndPassword(organization.Password, handlerForm["currentPassword"])
		if err != nil {
			HandleHTTPError(w, http.StatusUnauthorized, errors.New("Senha inválida"))
			return
		}

		newPassword := strings.TrimSpace(handlerForm["newPassword"])

		repo.ResetPasswordTo(organization, newPassword)

		HandleHTTPSuccess(w, nil)
	}
}

// ResetPasswordHandler resets the current user password, if it has the permissions
func ResetPasswordHandler(repo UpdatePasswordOrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var handlerForm map[string]string

		err := requestToJSONObject(r, &handlerForm)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		if HasPermission(r, model.PasswordResetPermission) == false {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("you do not have the right permissions to perform this action"))
			return
		}

		userID := GetUserID(r)
		organization, err := repo.Get(userID)

		newPassword := strings.TrimSpace(handlerForm["newPassword"])

		repo.ResetPasswordTo(organization, newPassword)

		HandleHTTPSuccess(w, nil)
	}
}

// ForgotPasswordHandler create a token to reset password and send it to email
func ForgotPasswordHandler(repo UpdatePasswordOrganizationRepository, mailer mail.Mailer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var handlerForm map[string]string

		err := requestToJSONObject(r, &handlerForm)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		organization, err := repo.GetByEmail(handlerForm["email"])
		if organization == nil {
			HandleHTTPError(w, http.StatusUnauthorized, errors.New("Email não encontrado"))
			os.Exit(1)
		}

		options := getJWTOptions()
		manager := JWTManager{OP: options}

		p := []string{model.PasswordResetPermission}
		token, err := manager.CreateToken(organization.User, &p)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		Emjson := mail.EmailParams{
			To:         organization.User.Email,
			Subject:    "Esqueci minha senha",
			Body:       token,
			TemplateID: "13f681b3-d48c-4b66-9204-e07f3afb33ed",
			Variables: map[string]string{
				"{{name}}": organization.Name,
				"{{link}}": token,
			},
		}

		err = mailer.SendEmail(Emjson)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao enviar o email"))
			os.Exit(1)
		}

		HandleHTTPSuccessNoContent(w)
	}
}

func getJWTOptions() JWTOptions {
	return JWTOptions{
		SigningMethod: os.Getenv("VITRINESOCIAL_SIGNING_METHOD"),
		PrivateKey:    []byte(os.Getenv("VITRINESOCIAL_PRIVATE_KEY")), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     []byte(os.Getenv("VITRINESOCIAL_PUBLIC_KEY")),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    24 * 3 * time.Hour,
	}
}
