package handlers

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Coderockr/vitrine-social/server/mail"
	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// UpdatePasswordOrganizationRepository represet operations for organization repository.
	UpdatePasswordOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
		GetByEmail(email string) (*model.Organization, error)
		ResetPasswordTo(o *model.Organization, password string) error
		ChangePassword(o model.Organization, cPassword, newPassword string) (model.Organization, error)
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

		var cPassword, nPassword string
		var ok bool

		if cPassword, ok = handlerForm["currentPassword"]; !ok {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("você deve informar senha atual"))
			return
		}

		if nPassword, ok = handlerForm["newPassword"]; !ok {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("você deve informar a nova senha"))
			return
		}

		if _, err = repo.ChangePassword(*organization, cPassword, nPassword); err != nil {
			HandleHTTPError(w, http.StatusUnauthorized, err)
			return
		}

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
func ForgotPasswordHandler(repo UpdatePasswordOrganizationRepository, mailer mail.Mailer, jm *JWTManager) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var handlerForm map[string]string

		err := requestToJSONObject(r, &handlerForm)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		organization, err := repo.GetByEmail(handlerForm["email"])
		if organization == nil {
			HandleHTTPError(w, http.StatusUnauthorized, errors.New("Não há organização com este email"))
			return
		}

		p := []string{model.PasswordResetPermission}
		token, err := jm.CreateToken(organization.User, &p)
		if err != nil {
			return
		}

		emailParams := mail.EmailParams{
			To:       organization.User.Email,
			Subject:  "Esqueci minha senha",
			Template: mail.ForgotPasswordTemplate,
			Variables: map[string]string{
				"name": organization.Name,
				"link": os.Getenv("FRONTEND_URL") + "/recover-password/" + token,
			},
		}

		if err := mailer.SendEmail(emailParams); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Falha ao enviar o email"))
			return
		}

		HandleHTTPSuccessNoContent(w)
	}
}
