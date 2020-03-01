package model

import "time"

// Token ...
type Token struct {
	ID           int       `json:"id"`
	UserId       int       `json:"userId"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	UserAgent    string    `json:"-"`
	Ip           string    `json:"-"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    string    `json:"createdAt"`
}
