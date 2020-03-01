package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	MiddleName        string `json:"middleName"`
	Image             string `json:"image"`
	IsActive          bool   `json:"isActive"`
	EncryptedPassword string `json:"-"`
	Password          string `json:"password,omitempty"`
	CreatorUserId     int    `json:"creatorUserId"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&u.Username, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 255)),
		validation.Field(&u.FirstName, validation.Length(0, 255)),
		validation.Field(&u.LastName, validation.Length(0, 255)),
		validation.Field(&u.MiddleName, validation.Length(0, 255)),
	)
}

// Before create ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
