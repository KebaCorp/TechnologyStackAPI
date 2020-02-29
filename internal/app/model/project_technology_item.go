package model

// ProjectTechnologyItem ...
type ProjectTechnologyItem struct {
	ID               int    `json:"id"`
	ProjectId        int    `json:"projectId"`
	TechnologyItemId int    `json:"technologyItemId"`
	StartedAt        string `json:"startedAt"`
	EndedAt          string `json:"endedAt"`
	CreatorUserId    int    `json:"creatorUserId"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}
