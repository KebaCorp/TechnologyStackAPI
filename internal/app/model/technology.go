package model

// Technology ...
type Technology struct {
	ID            int
	TypeId        int
	StageId       int
	Title         string
	IsDeprecated  bool
	CreatorUserId int
	CreatedAt     string
	UpdatedAt     string
}
