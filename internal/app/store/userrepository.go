package store

import (
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create user ...
func (r *UserRepository) CreateUser(u *model.User) (*model.User, error) {
	query := `INSERT INTO users (
		email,
		username,
		first_name,
		last_name,
		middle_name,
		image,
		is_active,
		encrypted_password,
		creator_user_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		u.Email,
		u.Username,
		u.FirstName,
		u.LastName,
		u.MiddleName,
		u.Image,
		u.IsActive,
		u.EncryptedPassword,
		u.CreatorUserId,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// Find by username or email ...
func (r *UserRepository) FindByUsernameOrEmail(username string, email string) (*model.User, error) {
	u := &model.User{}

	query := `SELECT
		id,
		email,
		username,
		first_name,
		last_name,
		middle_name,
		image,
		is_active,
		creator_user_id,
		created_at,
		updated_at
	FROM users
	WHERE username = $1 OR email = $2`

	if err := r.store.db.QueryRow(
		query,
		username,
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.MiddleName,
		&u.Image,
		&u.IsActive,
		&u.CreatorUserId,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
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
		image,
		is_active,
		creator_user_id,
		created_at,
		updated_at
	FROM users
	WHERE email = $1`

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
		&u.Image,
		&u.IsActive,
		&u.CreatorUserId,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return u, nil
}

// FindAll ...
func (r *UserRepository) FindAll() ([]*model.User, error) {
	query := `SELECT
		id,
		email,
		username,
		first_name,
		last_name,
		middle_name,
		image,
		is_active,
		creator_user_id,
		created_at,
		updated_at
	FROM users
	ORDER BY id ASC`

	rows, err := r.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*model.User, 0)

	for rows.Next() {
		u := new(model.User)
		if err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.Username,
			&u.FirstName,
			&u.LastName,
			&u.MiddleName,
			&u.Image,
			&u.IsActive,
			&u.CreatorUserId,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	return users, nil
}

// Delete user ...
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
