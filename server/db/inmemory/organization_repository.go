package inmemory

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
)

// OrganizationRepository is a implementation for tests
type OrganizationRepository struct {
	Storage map[string]model.Organization
}

// GetByEmail returns a organization by its email
func (r *OrganizationRepository) GetByEmail(email string) (*model.Organization, error) {
	for _, o := range r.Storage {
		if o.User.Email == email {
			return &o, nil
		}
	}

	return &model.Organization{}, errors.New("Login ou senha inválido")
}
