package workflow

import (
	"JStock/src/domain/models/repos"
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
	WorkFlowINfo *VWorkflowInfo `gorm:"embedded"`
	Repo               repos.IWorkFlowModel `gorm:"-"`
}

func NewWorkFlow(attrs ...WorkFlowAttrFunc) *WorkFlow {
	c := &WorkFlow{}
	WorkFlowAttrFuncs(attrs).apply(c)
	return c
}
