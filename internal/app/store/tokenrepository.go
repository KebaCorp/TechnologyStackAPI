package store

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"github.com/KebaCorp/TechnologyStackAPI/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

// TokenRepository ...
type TokenRepository struct {
	store *Store
}

// Update or create token ...
func (r *TokenRepository) UpdateOrCreateToken(
	user *model.User,
	userAgent string,
	ip string,
	expiresAt time.Time,
) (*model.Token, error) {
	query := `DELETE FROM tokens
	WHERE user_id = $1 AND user_agent = $2`

	if _, err := r.store.db.Exec(query, user.ID, userAgent); err != nil {
		return nil, err
	}

	token := &model.Token{
		UserId:       user.ID,
		Token:        GenerateToken(user.Username),
		RefreshToken: GenerateToken(user.Username),
		UserAgent:    userAgent,
		Ip:           ip,
		ExpiresAt:    expiresAt,
	}

	query = `INSERT INTO tokens (
		user_id,
		token,
		refresh_token,
		user_agent,
		ip,
		expires_at
	)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at`

	if err := r.store.db.QueryRow(
		query,
		token.UserId,
		token.Token,
		token.RefreshToken,
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

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)

	return hex.EncodeToString(hasher.Sum(nil))
}
