package handlers

import (
	"net/http"
	"strings"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	organizationRepository interface {
		Get(id int64) (*model.Organization, error)
		ResetPasswordTo(o *model.Organization, password string) error
	}
)

// UpdatePasswordHandler update user password
func UpdatePasswordHandler(repo organizationRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var handlerForm map[string]string

		err := requestToJSONObject(r, &handlerForm)
		if err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		id := GetUserID(r)
		organization, err := repo.Get(id)

		//currentPassword := handlerForm["currentPassword"]
		newPassword := strings.TrimSpace(handlerForm["newPassword"])

		repo.ResetPasswordTo(organization, newPassword)

		HandleHTTPSuccess(w, nil)
	}
}
