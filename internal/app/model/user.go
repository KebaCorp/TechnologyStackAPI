package model

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"firstName"`
	MiddleName        string `json:"firstName"`
	IsActive          bool   `json:"isActive"`
	EncryptedPassword string `json:"encryptedPassword"`
	CreatorUserId     int    `json:"creatorUserId"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
