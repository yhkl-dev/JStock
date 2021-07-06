package frontworkflowtempate

import (
	"JStock/src/domain/models/repos"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
)

type FrontWorkFlowItemTemplateAgg struct {
	WorkFlowItemTemplateMain     *workflowtemplate.WorkFlowItemTemplate
	WorkFlowItemTemplateMainRepo repos.IWorkFlowItemTemplateModel
}

func NewFrontWorkFlowItemTemplateAgg(m *workflowtemplate.WorkFlowItemTemplate, repo2 repos.IWorkFlowItemTemplateModel) *FrontWorkFlowItemTemplateAgg {
	if m == nil {
		panic("Error Root workFlowItemTempateMain")
	}
	fu := &FrontWorkFlowItemTemplateAgg{WorkFlowItemTemplateMain: m, WorkFlowItemTemplateMainRepo: repo2}
	fu.WorkFlowItemTemplateMain.Repo = repo2
	return fu
}

func (s *FrontWorkFlowItemTemplateAgg) CreateItem(m *workflowtemplate.WorkFlowItemTemplate) error {
	return s.WorkFlowItemTemplateMain.New()
}

func (s *FrontWorkFlowItemTemplateAgg) UpdateItem(m *workflowtemplate.WorkFlowItemTemplate) error {
	return s.WorkFlowItemTemplateMain.Update()
}
