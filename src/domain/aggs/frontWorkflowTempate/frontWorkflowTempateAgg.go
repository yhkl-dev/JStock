package frontworkflowtempate

import (
	"JStock/src/domain/models/repos"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
	"JStock/src/infrastructure/Errors"
)

type FrontWorkFlowTemplateAgg struct {
	WorkFlowTempateMain          *workflowtemplate.WorkFlowTemplate
	WorkFlowItemTemplateMain     *workflowtemplate.WorkFlowItemTemplate
	WorkFlowTemplateMainRepo     repos.IWorkFlowTemplateModel
	WorkFlowItemTemplateMainRepo repos.IWorkFlowItemTemplateModel
}

func NewFrontWorkFlowTemplateAgg(workFlowTempateMain *workflowtemplate.WorkFlowTemplate, itemMain *workflowtemplate.WorkFlowItemTemplate, repo repos.IWorkFlowTemplateModel, repo2 repos.IWorkFlowItemTemplateModel) *FrontWorkFlowTemplateAgg {
	if workFlowTempateMain == nil {
		panic("Error Root workFlowTempateMain")
	}
	fu := &FrontWorkFlowTemplateAgg{WorkFlowTempateMain: workFlowTempateMain, WorkFlowItemTemplateMain: itemMain, WorkFlowTemplateMainRepo: repo, WorkFlowItemTemplateMainRepo: repo2}
	fu.WorkFlowTempateMain.Repo = repo
	fu.WorkFlowItemTemplateMain.Repo = repo2
	return fu
}

func (s *FrontWorkFlowTemplateAgg) QueryDetail() error {
	if s.WorkFlowTempateMain.ID <= 0 {
		return Errors.NewNotFoundIDError("WorkFlowTempateMain")
	}
	err := s.WorkFlowTempateMain.Load()
	if err != nil {
		return Errors.NewNotFoundDataError("WorkFlowTempateMain", err.Error())
	}
	results, err := s.WorkFlowItemTemplateMain.List(s.WorkFlowTempateMain.ID)
	if err != nil {
		return Errors.NewNotFoundDataError("WorkFlowTempateMain", err.Error())
	}
	s.WorkFlowTempateMain.FlowItems = results
	return nil
}

func (s *FrontWorkFlowTemplateAgg) CreateFlowTemplate(m *workflowtemplate.WorkFlowTemplate) error {
	return s.WorkFlowTempateMain.New()
}

func (s *FrontWorkFlowTemplateAgg) ListFlowTemplate(m *workflowtemplate.WorkFlowTemplate, page, pageSize int) (interface{}, error) {
	results, err := s.WorkFlowTempateMain.ListTemplate(m.FlowName, m.FlowType, page, pageSize)
	if err != nil {
		return nil, err
	}
	for _, m := range results.([]*workflowtemplate.WorkFlowTemplate) {

		res, err := s.WorkFlowItemTemplateMain.List(m.ID)
		if err != nil {
			return nil, Errors.NewNotFoundDataError("UserRoleMap", err.Error())
		}
		m.FlowItems = res
	}
	return results, nil
}
