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

// Create technology ...
func (r *TechnologyRepository) CreateTechnology(t *model.Technology) (int, error) {
	query := `INSERT INTO
	technologies (
		type_id,
		stage_id,
		title,
		image,
		is_deprecated,
		creator_user_id
	)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		t.TypeId,
		t.StageId,
		t.Title,
		t.Image,
		t.IsDeprecated,
		t.CreatorUserId,
	).Scan(&t.ID); err != nil {
		return 0, err
	}

	return t.ID, nil
}

// Delete technology ...
func (r *TechnologyRepository) DeleteTechnology(id int) error {
	query := `DELETE FROM technologies WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
