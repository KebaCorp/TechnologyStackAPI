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
	rows, err := r.store.db.Query("SELECT * FROM types WHERE is_deleted IS NOT true ORDER BY id ASC")

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
			&t.IsDeleted,
			&t.CreatorUserId,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}

		types = append(types, t)
	}

	return types, nil
}

// Create type ...
func (r *TypeRepository) CreateType(t *model.Type) (int, error) {
	query := `INSERT INTO
	 types (title, is_deleted, creator_user_id)
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
	query := `DELETE FROM types WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
