package store

import (
	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
)

// TokenRepository ...
type TokenRepository struct {
	store *Store
}

// Update or create token ...
func (r *TokenRepository) UpdateOrCreateToken(token *model.Token) (*model.Token, error) {
	query := `DELETE FROM tokens
	WHERE user_id = $1 AND user_agent = $2`

	if _, err := r.store.db.Exec(query, token.UserId, token.UserAgent); err != nil {
		return nil, err
	}

	query = `INSERT INTO tokens (
		user_id,
		user_agent,
		ip,
		expires_at
	)
	VALUES($1, $2, $3, $4)
	RETURNING id, created_at`

	if err := r.store.db.QueryRow(
		query,
		token.UserId,
		token.UserAgent,
		token.Ip,
		token.ExpiresAt,
	).Scan(
		&token.ID,
		&token.CreatedAt,
	); err != nil {
		return nil, err
	}

	return token, nil
}
