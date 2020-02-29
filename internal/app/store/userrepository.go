package store

import "github.com/KebaCorp/TechnologyStackAPI/internal/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create user ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	query := `INSERT INTO users (
		email,
		username,
		first_name,
		last_name,
		middle_name,
		is_active,
		encrypted_password,
		creator_user_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		u.Email,
		u.Username,
		u.FirstName,
		u.LastName,
		u.MiddleName,
		u.IsActive,
		u.EncryptedPassword,
		u.CreatorUserId,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	query := `SELECT
		id,
		email,
		username,
		first_name,
		last_name,
		middle_name,
		is_active,
		encrypted_password,
		creator_user_id
	FROM users WHERE email = $1`

	if err := r.store.db.QueryRow(
		query,
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.MiddleName,
		&u.IsActive,
		&u.EncryptedPassword,
		&u.CreatorUserId,
	); err != nil {
		return nil, err
	}

	return u, nil
}
