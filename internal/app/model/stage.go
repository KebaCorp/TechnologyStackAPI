package model

// Stage ...
type Stage struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	IsDeleted     bool          `json:"isDeleted"`
	CreatorUserId int           `json:"creatorUserId"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	Technologies  []*Technology `json:"technologies"`
}
