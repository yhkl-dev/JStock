package GormDao

import (
	"JStock/src/domain/models/repos"
	rolemodel "JStock/src/domain/models/roleModel"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"

	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
)

type ItemMainRepo struct {
	db *gorm.DB
}

func NewItemMainRepo(db *gorm.DB) *ItemMainRepo {
	return &ItemMainRepo{db: db}
}

func (s *ItemMainRepo) New(model repos.IModel) error {
	m := model.(*workflowtemplate.WorkFlowItemTemplate)
	sql, args, _ := squirrel.Insert("t_workflow_item_template").
		Columns("template_id", "item_name", "exec_order", "role_id").
		Values(m.TemplateID, m.ItemName, m.ExecOrder, m.RoleID).
		ToSql()
	err := s.db.Exec(sql, args...).Error
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
	id := model.(*workflowtemplate.WorkFlowItemTemplate).ID
	sql, args, _ := squirrel.
		Select("twit.id, twit.template_id ,twt.flow_name template_name,twit.exec_order exec_order,twit.role_id,tr.role_name,twit.item_name").
		From("t_workflow_item_template twit,t_workflow_template twt,t_roles tr").
		Where("twt.id = twit.template_id and twit.role_id = tr.id and twit.id = ?", id).ToSql()
	return  s.db.Raw(sql, args...).Scan(model).Error
}

func (s *ItemMainRepo) List(id int) (interface{}, error) {
	var res = make([]*workflowtemplate.WorkFlowItemTemplate, 0)
	sql, args, _ := squirrel.
		Select("twit.id, twit.template_id ,twt.flow_name template_name,twit.exec_order exec_order,twit.role_id,tr.role_name,twit.item_name").
		From("t_workflow_item_template twit,t_workflow_template twt,t_roles tr").
		Where("twt.id = twit.template_id and twit.role_id = tr.id and twt.id = ?", id).
		OrderBy("twit.id").ToSql()
	err := s.db.Raw(sql, args...).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *ItemMainRepo) Update(model repos.IModel) error {
	m := model.(*workflowtemplate.WorkFlowItemTemplate)
	updateData := map[string]interface{}{
		"template_id": m.TemplateID,
		"item_name":  m.ItemName,
		"exec_order": m.ExecOrder,
		"role_id": m.RoleID,
	}
	sql, args, _ := squirrel.Update("t_workflow_item_template").SetMap(updateData).Where("id =?", m.ID).ToSql()
	return s.db.Exec(sql, args...).Error
}