package model

// Stage ...
type Stage struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	CreatorUserId int           `json:"creatorUserId"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	Technologies  []*Technology `json:"technologies"`
	IsDeleted     bool          `json:"isDeleted"`
}
