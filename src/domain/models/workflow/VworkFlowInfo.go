package workflow

import "time"

type VWorkflowInfo struct {
	WorkFlowTemplateID int       `josn:"workflow_template_id" gorm:"column:workflow_template_id"`
	MNCTemplateID      int       `json:"mnc_template_id" gorm:"column:mnc_template_id"`
	CreateUserID       int       `json:"user_id" gorm:"column:user_id"`
	IsDeleted          int       `json:"is_deleted" gorm:"column:is_deleted"`
	Status             int       `json:"status" gorm:"column:status"`
	CreateAt           time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt           time.Time `json:"update_at" gorm:"column:update_at"`
}

func NewVWorkflowInfo() *VWorkflowInfo {
	return &VWorkflowInfo{}
}