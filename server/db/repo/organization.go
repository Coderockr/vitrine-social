package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Coderockr/vitrine-social/server/security"
	"github.com/gobuffalo/pop/nulls"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/jmoiron/sqlx"
)

// OrganizationRepository is a implementation for Postgres
type OrganizationRepository struct {
	db      *sqlx.DB
	catRepo *CategoryRepository
}

// NewOrganizationRepository creates a new repository
func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{
		db:      db,
		catRepo: NewCategoryRepository(db),
	}
}

const allFields string = `
	id, name, logo_image_id, phone, about, video, email, password, slug,
	street as "address.street", number as "address.number",
	complement as "address.complement", neighborhood as "address.neighborhood",
	city as "address.city", state as "address.state", zipcode as "address.zipcode",
	website
`

// GetBaseOrganization returns only the data about a organization, not its relations
func (r *OrganizationRepository) GetBaseOrganization(id int64) (*model.Organization, error) {
	o := &model.Organization{}
	err := r.db.Get(o, "SELECT "+allFields+" FROM organizations WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	o.Logo, err = r.GetLogo(*o)

	return o, err
}

// Get one Organization from database
func (r *OrganizationRepository) Get(id int64) (*model.Organization, error) {
	o, err := r.GetBaseOrganization(id)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&o.Images, "SELECT * FROM organizations_images WHERE organization_id = $1 AND id != $2", id, o.LogoImageID)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&o.Needs, "SELECT * FROM needs WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	for i := range o.Needs {
		o.Needs[i].Category, err = r.catRepo.Get(o.Needs[i].CategoryID)
		if err != nil {
			fmt.Println("test?")
			return nil, err
		}

		o.Needs[i].Images, err = getNeedImages(r.db, &o.Needs[i])
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

// Create receives a Organization and creates it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Create(o model.Organization) (model.Organization, error) {
	row := r.db.QueryRow(
		`INSERT INTO organizations (
			name, phone, about, video, email, slug, password,
			street, number, complement, neighborhood, city, state, zipcode,
			website
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		)
		RETURNING id
		`,
		o.Name,
		o.Phone,
		o.About,
		o.Video,
		o.Email,
		o.Slug,
		o.Password,
		o.Address.Street,
		o.Address.Number,
		o.Address.Complement,
		o.Address.Neighborhood,
		o.Address.City,
		o.Address.State,
		o.Address.Zipcode,
		o.Website,
	)

	err := row.Scan(&o.ID)

	if err != nil {
		return o, err
	}

	return o, nil
}

// Update - Receive an Organization and update it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Update(o model.Organization) (model.Organization, error) {
	_, err := r.db.Exec(
		`UPDATE organizations SET
			name = $1,
			phone = $2,
			about = $3,
			video = $4,
			email = $5,
			street = $6,
			number = $7,
			complement = $8,
			neighborhood = $9,
			city = $10,
			state = $11,
			zipcode = $12,
			website = $13
		WHERE id = $14
		`,
		o.Name,
		o.Phone,
		o.About,
		o.Video,
		o.Email,
		o.Address.Street,
		o.Address.Number,
		o.Address.Complement,
		o.Address.Neighborhood,
		o.Address.City,
		o.Address.State,
		o.Address.Zipcode,
		o.Website,
		o.ID,
	)

	if err != nil {
		return o, err
	}

	return o, nil
}

// DeleteImage - Receive an id and remove the image
func (r *OrganizationRepository) DeleteImage(imageID int64, organizationID int64) error {
	_, err := r.db.Exec(`DELETE FROM organizations_images WHERE id = $1 AND organization_id = $2`, imageID, organizationID)
	return err
}

// GetByEmail returns a organization by its email
func (r *OrganizationRepository) GetByEmail(email string) (*model.Organization, error) {
	o := model.Organization{}
	err := r.db.Get(&o, `SELECT `+allFields+` FROM organizations WHERE email = $1`, email)

	if err != nil {
		return nil, err
	}

	o.Logo, err = r.GetLogo(o)

	return &o, err
}

// GetUserByEmail returns a organization user by its email
func (r *OrganizationRepository) GetUserByEmail(email string) (model.User, error) {
	o, err := r.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	return o.User, nil
}

// ResetPasswordTo resets the organization password to the value informed
func (r *OrganizationRepository) ResetPasswordTo(o *model.Organization, password string) error {
	hash, err := security.GenerateHash(password)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`UPDATE organizations SET password = $1 WHERE id = $2`, hash, o.ID)
	if err != nil {
		return err
	}
	o.Password = hash
	return nil
}

// CreateImage creates a new organization image based on the struct
func (r *OrganizationRepository) CreateImage(i model.OrganizationImage) (model.OrganizationImage, error) {
	err := r.db.QueryRow(
		`INSERT INTO organizations_images (organization_id, name, url)
			VALUES($1, $2, $3)
			RETURNING id
		`,
		i.OrganizationID,
		i.Image.Name,
		i.Image.URL,
	).Scan(&i.ID)

	if err != nil {
		return i, err
	}

	return i, nil
}

// ChangePassword will update a organization password given its old password
func (r *OrganizationRepository) ChangePassword(o model.Organization, currentPassword, newPassword string) (model.Organization, error) {
	err := security.CompareHashAndPassword(o.Password, currentPassword)
	if err != nil {
		return o, errors.New("Senha inválida")
	}

	newPassword = strings.TrimSpace(newPassword)

	r.ResetPasswordTo(&o, newPassword)
	return o, nil
}

// UpdateLogo will change the logo image
func (r *OrganizationRepository) UpdateLogo(imageID nulls.Int64, organizationID int64) error {
	_, err := r.db.Exec(`UPDATE organizations SET logo_image_id = $1 WHERE id = $2`, imageID, organizationID)
	return err
}

// GetLogo returns organization logo image
func (r *OrganizationRepository) GetLogo(o model.Organization) (*model.OrganizationImage, error) {
	logo := &model.OrganizationImage{}

	logoID, err := o.LogoImageID.Value()
	if logoID == nil || err != nil {
		return logo, nil
	}

	err = r.db.Get(logo, "SELECT * FROM organizations_images WHERE id = $1", o.LogoImageID)
	if err == sql.ErrNoRows {
		return logo, nil
	}

	return logo, err
}
