package store

import (
	"context"
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

var (
	ctx context.Context
)

// TechnologyRepository ...
type TechnologyRepository struct {
	store *Store
}

// FindAll ...
func (r *TechnologyRepository) FindAll() ([]*model.Technology, error) {
	rows, err := r.store.db.Query("SELECT * FROM technologies ORDER BY type_id ASC, stage_id ASC")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	technologies := make([]*model.Technology, 0)

	for rows.Next() {
		t := new(model.Technology)
		if err := rows.Scan(
			&t.ID,
			&t.TypeId,
			&t.StageId,
			&t.Title,
			&t.Image,
			&t.IsDeprecated,
			&t.CreatorUserId,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}

		technologies = append(technologies, t)
	}

	return technologies, nil
}
