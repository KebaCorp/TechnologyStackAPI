package model

// Token ...
type Token struct {
	ID        int    `json:"id"`
	UserId    int    `json:"userId"`
	UserAgent string `json:"userAgent"`
	Ip        string `json:"ip"`
	ExpiresAt string `json:"expiresAt"`
	CreatedAt string `json:"createdAt"`
}
