package importanctLevelModel

import "time"

type ImportancyLevel struct {
	ID        int       `json:"id" gorm:"id"`
	LevelName string    `json:"level_name" gorm:"level_name"`
	Comment   string    `json:"comment" gorm:"comment"`
	CreateAt  time.Time `json:"create_at" gorm:"create_at"`
}
