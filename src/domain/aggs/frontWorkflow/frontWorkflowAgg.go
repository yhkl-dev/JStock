package frontworkflow

import (
	"JStock/src/domain/models/repos"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
	"JStock/src/domain/models/workflow"
)

type FrontWorkFlowAgg struct {
	WorkFlowMain *workflow.WorkFlow
	WorkFlowTempateMain          *workflowtemplate.WorkFlowTemplate
	WorkFlowMainRepo repos.IWorkFlowModel
	WorkFlowTemplateMainRepo     repos.IWorkFlowTemplateModel
}

func NewFrontWorkFlowAgg(wf *workflow.WorkFlow, wft *workflowtemplate.WorkFlowTemplate, repo1 repos.IWorkFlowModel, repo2 repos.IWorkFlowTemplateModel) *FrontWorkFlowAgg{
	if wf == nil {
		panic("Error Root WorkFlowMain") 
	}
	f := &FrontWorkFlowAgg{WorkFlowMain: wf, WorkFlowTempateMain:  wft, WorkFlowMainRepo: repo1, WorkFlowTemplateMainRepo: repo2}
	f.WorkFlowMain.Repo = repo1
	f.WorkFlowTempateMain.Repo = repo2
	return f
}

func (s *FrontWorkFlowAgg) CreateWorkFlow(m *workflow.WorkFlow) error {
	return s.WorkFlowTempateMain.New()
}