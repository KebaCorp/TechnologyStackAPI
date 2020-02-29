package store

import (
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// TechnologyItemRepository ...
type TechnologyItemRepository struct {
	store *Store
}

// FindAll ...
func (r *TechnologyItemRepository) FindAll() ([]*model.TechnologyItem, error) {
	query := `SELECT
		id,
		technology_id,
		parent_id,
		title,
		description,
		creator_user_id,
		created_at,
		updated_at,
		is_deleted
	FROM technology_items
	WHERE is_deleted IS NOT true
	ORDER BY id ASC`

	rows, err := r.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	technologyItems := make([]*model.TechnologyItem, 0)

	for rows.Next() {
		t := new(model.TechnologyItem)
		if err := rows.Scan(
			&t.ID,
			&t.TechnologyId,
			&t.ParentId,
			&t.Title,
			&t.Description,
			&t.CreatorUserId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.IsDeleted,
		); err != nil {
			log.Fatal(err)
		}

		technologyItems = append(technologyItems, t)
	}

	return technologyItems, nil
}

// Create technology item ...
func (r *TechnologyItemRepository) CreateTechnologyItem(t *model.TechnologyItem) (int, error) {
	query := `INSERT INTO
	technology_items (
		technology_id,
		parent_id,
		title,
		description,
		creator_user_id
	)
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		&t.TechnologyId,
		&t.ParentId,
		&t.Title,
		&t.Description,
		&t.CreatorUserId,
	).Scan(&t.ID); err != nil {
		return 0, err
	}

	return t.ID, nil
}

// Delete technology item ...
func (r *TechnologyItemRepository) DeleteTechnologyItem(id int) error {
	query := `UPDATE technology_items SET is_deleted = true WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
