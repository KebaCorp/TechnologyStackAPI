package model

// Type ...
type Type struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	CreatorUserId int           `json:"creatorUserId"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	Stages        []*Stage      `json:"stages"`
	Technologies  []*Technology `json:"technologies"`
	IsDeleted     bool          `json:"isDeleted"`
}
