package GormDao

import (
	"JStock/src/domain/models/repos"
	rolemodel "JStock/src/domain/models/roleModel"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"

	"gorm.io/gorm"
)

type ItemMainRepo struct {
	db *gorm.DB
}

func NewItemMainRepo(db *gorm.DB) *ItemMainRepo {
	return &ItemMainRepo{db: db}
}

func (s *ItemMainRepo) New(model repos.IModel) error {
	err := s.db.Table("t_workflow_item_template").Create(model).Error
	if err != nil {
		return err
	}
	role := rolemodel.New()
	err = s.db.Table("t_roles").Where("id = ?", model.(*workflowtemplate.WorkFlowItemTemplate).RoleID).Find(role).Error
	if err != nil {
		return err
	}
	model.(*workflowtemplate.WorkFlowItemTemplate).RoleName = role.RoleInfo.RoleName
	workflow := workflowtemplate.NewWorkFlowTemplate()
	err = s.db.Table("t_workflow_template").Where("id = ?", model.(*workflowtemplate.WorkFlowItemTemplate).TemplateID).Find(workflow).Error
	if err != nil {
		return err
	}
	model.(*workflowtemplate.WorkFlowItemTemplate).TemplateName = workflow.FlowName
	return nil
}

func (s *ItemMainRepo) FindByID(model repos.IModel) error {
	return s.db.Table("t_workflow_item_template").Where("id = ? ", model.(*workflowtemplate.WorkFlowItemTemplate).ID).First(model).Error
}

func (s *ItemMainRepo) List(id int) (interface{}, error) {
	var res = make([]*workflowtemplate.WorkFlowItemTemplate, 0)
	err := s.db.Table("t_workflow_item_template").Find(&res).Order("exec_order").Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
