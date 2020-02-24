package model

import "database/sql"

// Technology ...
type Technology struct {
	ID            int            `json:"id"`
	TypeId        int            `json:"typeId"`
	StageId       int            `json:"stageId"`
	Title         string         `json:"title"`
	Image         sql.NullString `json:"image"`
	IsDeprecated  bool           `json:"isDeprecated"`
	CreatorUserId int            `json:"creatorUserId"`
	CreatedAt     string         `json:"createdAt"`
	UpdatedAt     string         `json:"updatedAt"`
}
