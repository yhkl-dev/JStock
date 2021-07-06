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
	var m = make([]*workflowtemplate.WorkFlowTemplate, 0)
	s.db = s.db.Table("t_workflow_template")
	if flowname != "" {
		s.db = s.db.Where("flow_name = ?", flowname)
	}
	if flowType != 0 {
		s.db = s.db.Where("flow_type = ?", flowType)
	}
	if pageSize == 0 {
		pageSize = 10
	}
	err := s.db.Find(&m).Offset(page).Limit(pageSize).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}


func (s *FlowTemplateRepo) Update(model repos.IModel) error {
	return s.db.Table("t_workflow_template").Where("id = ? ", model.(*workflowtemplate.WorkFlowTemplate).ID).First(model).Error
}