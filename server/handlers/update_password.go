package handlers

import (
	"errors"
	"net/http"

	"github.com/Coderockr/vitrine-social/server/model"
)

type (
	// UpdatePasswordOrganizationRepository represet operations for organization repository.
	UpdatePasswordOrganizationRepository interface {
		Get(id int64) (*model.Organization, error)
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
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Você deve informar senha atual !"))
			return
		}

		if nPassword, ok = handlerForm["newPassword"]; !ok {
			HandleHTTPError(w, http.StatusBadRequest, errors.New("Você deve informar a nova senha !"))
			return
		}

		if _, err = repo.ChangePassword(*organization, cPassword, nPassword); err != nil {
			HandleHTTPError(w, http.StatusBadRequest, err)
			return
		}

		HandleHTTPSuccess(w, nil)
	}
}
