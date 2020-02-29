package store

import (
	"log"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// ProjectRepository ...
type ProjectRepository struct {
	store *Store
}

// FindAll ...
func (r *ProjectRepository) FindAll() ([]*model.Project, error) {
	query := `SELECT
		id,
		title,
		code,
		image,
		is_active,
		creator_user_id,
		created_at,
		updated_at
	FROM projects
	ORDER BY id ASC`

	rows, err := r.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := make([]*model.Project, 0)

	for rows.Next() {
		t := new(model.Project)
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Code,
			&t.Image,
			&t.IsActive,
			&t.CreatorUserId,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}

		projects = append(projects, t)
	}

	return projects, nil
}

// Create project ...
func (r *ProjectRepository) CreateProject(t *model.Project) (int, error) {
	query := `INSERT INTO
	 projects (title, code, image, is_active, creator_user_id)
	 VALUES($1, $2, $3, $4, $5) RETURNING id`

	if err := r.store.db.QueryRow(
		query,
		t.Title,
		t.Code,
		t.Image,
		t.IsActive,
		t.CreatorUserId,
	).Scan(&t.ID); err != nil {
		return 0, err
	}

	return t.ID, nil
}

// Delete project ...
func (r *ProjectRepository) DeleteProject(id int) error {
	query := `DELETE FROM projects WHERE id = $1`

	if _, err := r.store.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
