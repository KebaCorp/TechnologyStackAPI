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
