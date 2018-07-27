package repo

import (
	"database/sql"
	"errors"
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

	err = r.db.Select(&o.Images, "SELECT * FROM organizations_images WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&o.Needs, "SELECT * FROM needs WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}

	for i := range o.Needs {
		c, _ := r.catRepo.Get(o.Needs[i].CategoryID)
		o.Needs[i].Category = *c

		o.Needs[i].Images, err = getNeedImages(r.db, &o.Needs[i])
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

// Create receives a Organization and creates it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Create(o model.Organization) (model.Organization, error) {
	o, err := validateOrg(o)
	if err != nil {
		return o, err
	}

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

	err = row.Scan(&o.ID)
	if err != nil {
		return o, err
	}

	return o, nil
}

func validateOrg(o model.Organization) (model.Organization, error) {

	o.Name = strings.TrimSpace(o.Name)
	o.Phone = strings.TrimSpace(o.Phone)
	o.About = strings.TrimSpace(o.About)
	o.Video = strings.TrimSpace(o.Video)
	o.Email = strings.TrimSpace(o.Email)
	o.Address.Street = strings.TrimSpace(o.Address.Street)
	o.Address.Number = strings.TrimSpace(o.Address.Number)
	o.Address.Neighborhood = strings.TrimSpace(o.Address.Neighborhood)
	o.Address.City = strings.TrimSpace(o.Address.City)
	o.Address.State = strings.TrimSpace(o.Address.State)
	o.Address.Zipcode = strings.TrimSpace(o.Address.Zipcode)
	if o.Address.Complement != nil {
		*o.Address.Complement = strings.TrimSpace(*o.Address.Complement)
	}

	if len(o.Name) == 0 {
		return o, errors.New("organization name should not be empty")
	}

	if len(o.Phone) == 0 {
		return o, errors.New("organization phone should not be empty")
	}

	if len(o.About) == 0 {
		return o, errors.New("organization about should not be empty")
	}

	if len(o.Email) == 0 {
		return o, errors.New("organization email should not be empty")
	}

	if len(o.Address.Street) == 0 {
		return o, errors.New("organization address street should not be empty")
	}

	if len(o.Address.Number) == 0 {
		return o, errors.New("organization address number should not be empty")
	}

	if len(o.Address.Neighborhood) == 0 {
		return o, errors.New("organization address neighborhood should not be empty")
	}

	if len(o.Address.City) == 0 {
		return o, errors.New("organization address city should not be empty")
	}

	if len(o.Address.State) == 0 {
		return o, errors.New("organization address state should not be empty")
	}

	if len(o.Address.Zipcode) == 0 {
		return o, errors.New("organization address zipcode should not be empty")
	}

	return o, nil
}

// Update - Receive an Organization and update it in the database, returning the updated Organization or error if failed
func (r *OrganizationRepository) Update(o model.Organization) (model.Organization, error) {
	o, err := validateOrg(o)
	if err != nil {
		return o, err
	}

	_, err = r.db.Exec(
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
