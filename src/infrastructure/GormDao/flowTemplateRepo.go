package GormDao

import (
	"JStock/src/domain/models/repos"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"

	"gorm.io/gorm"
)

type FlowTemplateRepo struct {
	db *gorm.DB
}

func NewIFlowTemplateRepo(db *gorm.DB) *FlowTemplateRepo {
	return &FlowTemplateRepo{db: db}
}
func (s *FlowTemplateRepo) New(model repos.IModel) error {
	return s.db.Table("t_workflow_template").Create(model).Error
}

func (s *FlowTemplateRepo) NewWorkFlowTemplte(model repos.IModel) error {
	return s.db.Table("t_workflow_template").Create(model).Error
}

func (s *FlowTemplateRepo) FindByID(model repos.IModel) error {
	return s.db.Table("t_workflow_template").Where("id = ? ", model.(*workflowtemplate.WorkFlowTemplate).ID).First(model).Error
}

func (s *FlowTemplateRepo) ListTemplate(flowname string, flowType, page, pageSize int) (interface{}, error) {
	var res = make([]*workflowtemplate.WorkFlowItemTemplate, 0)
	err := s.db.Table("t_workflow_template").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
