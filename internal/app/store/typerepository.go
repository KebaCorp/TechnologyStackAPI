package store

import (
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// TypeRepository ...
type TypeRepository struct {
	store *Store
}

// FindAll ...
func (r *TypeRepository) FindAll() ([]*model.Type, error) {
	query := `SELECT
		id,
		title,
		creator_user_id,
		created_at,
		updated_at,
		is_deleted
	FROM types
	WHERE is_deleted IS NOT true
	ORDER BY id ASC`

	rows, err := r.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	types := make([]*model.Type, 0)

	for rows.Next() {
		t := new(model.Type)
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.CreatorUserId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.IsDeleted,
		); err != nil {
			log.Fatal(err)
		}

		types = append(types, t)
	}

	return types, nil
}

// Create type ...
func (r *TypeRepository) CreateType(t *model.Type) (int, error) {
	query := `INSERT INTO types (title, is_deleted, creator_user_id)
	VALUES($1, $2, $3) RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		t.Title,
		t.IsDeleted,
		t.CreatorUserId,
	).Scan(&t.ID); err != nil {
		return 0, err
	}

	return t.ID, nil
}

// Delete type ...
func (r *TypeRepository) DeleteType(id int) error {
	query := `UPDATE types SET is_deleted = true WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
