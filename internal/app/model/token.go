package model

import "time"

// Token ...
type Token struct {
	ID        int       `json:"id"`
	UserId    int       `json:"userId"`
	UserAgent string    `json:"userAgent"`
	Ip        string    `json:"ip"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt string    `json:"createdAt"`
}
