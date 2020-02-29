package model

// Project ...
type Project struct {
	ID            int    `json:"id"`
	Title         int    `json:"title"`
	Code          int    `json:"code"`
	IsActive      bool   `json:"isActive"`
	CreatorUserId int    `json:"creatorUserId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
