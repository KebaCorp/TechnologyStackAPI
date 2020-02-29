package model

// Project ...
type Project struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Code          string `json:"code"`
	Image         string `json:"image"`
	IsActive      bool   `json:"isActive"`
	CreatorUserId int    `json:"creatorUserId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}
