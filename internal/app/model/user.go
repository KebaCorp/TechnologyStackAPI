package model

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"firstName"`
	MiddleName        string `json:"firstName"`
	Image             string `json:"image"`
	IsActive          bool   `json:"isActive"`
	EncryptedPassword string `json:"encryptedPassword"`
	Password          string `json:"password"`
	CreatorUserId     int    `json:"creatorUserId"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
