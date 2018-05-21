package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/Coderockr/vitrine-social/server/security"
)

type (
	// OrganizationRepository represet operations for organization repository.
	OrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
		ResetPasswordTo(o *model.Organization, password string) error
	}
)

// UpdatePasswordHandler update user password
func UpdatePasswordHandler(repo OrganizationRepository) func(w http.ResponseWriter, r *http.Request) {
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
