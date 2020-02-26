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
	rows, err := r.store.db.Query("SELECT * FROM stages WHERE is_deleted IS NOT true ORDER BY id ASC")

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

// Create stage ...
func (r *StageRepository) CreateStage(t *model.Stage) (int, error) {
	query := `INSERT INTO
	 stages (title, is_deleted, creator_user_id)
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

// Delete stage ...
func (r *StageRepository) DeleteStage(id int) error {
	query := `UPDATE stages SET is_deleted = true WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
