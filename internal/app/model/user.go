package model

// User ...
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
	CreatorUserId     int
	CreatedAt         string
	UpdatedAt         string
}
