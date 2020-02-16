package store

import (
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// StageRepository ...
type StageRepository struct {
	store *Store
}

// FindAll ...
func (r *StageRepository) FindAll() ([]*model.Stage, error) {
	rows, err := r.store.db.Query("SELECT * FROM stages WHERE is_deleted IS NOT true")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stages := make([]*model.Stage, 0)

	for rows.Next() {
		t := new(model.Stage)
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

		stages = append(stages, t)
	}

	return stages, nil
}
