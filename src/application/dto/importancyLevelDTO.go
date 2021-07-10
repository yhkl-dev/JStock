package dto

import "time"

type (
	ImportancyLevelMainResponse struct {
		ID        int       `json:"id" `
		LevelName string    `json:"level_name" `
		Comment   string    `json:"comment" `
		CreateAt  time.Time `json:"create_at" `
	}
	ImportancyLevelListResponse []ImportancyLevelMainResponse
)
