package model

// UserTechnologyItem ...
type UserTechnologyItem struct {
	ID               int    `json:"id"`
	TechnologyItemId int    `json:"technologyItemId"`
	UserId           int    `json:"userId"`
	StartedAt        string `json:"startedAt"`
	CreatorUserId    int    `json:"creatorUserId"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
