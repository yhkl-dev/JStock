package workflow

import (
	"JStock/src/domain/models/repos"
	"time"
)

type WorkFlowAttrFunc func(model *WorkFlow)
type WorkFlowAttrFuncs []WorkFlowAttrFunc

func (s WorkFlowAttrFuncs) apply(u *WorkFlow) {
	for _, f := range s {
		f(u)
	}
}

type WorkFlow struct {
	ID                 int             `json:"id" gorm:"column:id"`
	WorkFlowTemplateID int             `josn:"workflow_template_id" gorm:"column:workflow_template_id"`
	MNCTemplateID      int             `json:"mnc_template_id" gorm:"column:mnc_template_id"`
	CreateUserID       int             `json:"user_id" gorm:"column:user_id"`
	CreateAt           time.Time       `json:"create_at" gorm:"column:create_at"`
	UpdateAt           time.Time       `json:"update_at" gorm:"column:update_at"`
	IsDeleted          int             `json:"is_deleted" gorm:"column:is_deleted"`
	Status             int             `json:"status" gorm:"column:status"`
	Repo               repos.IWorkFlow `gorm:"-"`
}

func NewWorkFlow(attrs ...WorkFlowAttrFunc) *WorkFlow {
	c := &WorkFlow{}
	WorkFlowAttrFuncs(attrs).apply(c)
	return c
}
