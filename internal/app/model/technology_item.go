package model

// TechnologyItem ...
type TechnologyItem struct {
	ID            int    `json:"id"`
	TechnologyId  int    `json:"technologyId"`
	ParentId      int    `json:"parentId"`
	Title         int    `json:"title"`
	Description   string `json:"description"`
	CreatorUserId int    `json:"creatorUserId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	IsDeleted     bool   `json:"isDeleted"`
}
