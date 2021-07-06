package assembler

import (
	"JStock/src/application/dto"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
)

type ItemMainRequest struct{}

func (s *ItemMainRequest) D2M_ItemMain(dto *dto.ItemAddRequest) *workflowtemplate.WorkFlowItemTemplate {
	m := workflowtemplate.NewWorkFlowItemTempate()
	m.ExecOrder = dto.ExecOrder
	m.ItemName = dto.ItemName
	m.RoleID = dto.RoleID
	m.TemplateID = dto.TemplateID
	return m
}

type ItemUpdateRequest struct{}

func (s *ItemUpdateRequest) D2M_Item(dto *dto.ItemUpdateRequest) *workflowtemplate.WorkFlowItemTemplate {
	m := workflowtemplate.NewWorkFlowItemTempate()
	m.ID = dto.ID
	m.RoleID = dto.RoleID
	m.TemplateID = dto.TemplateID
	m.ItemName = dto.ItemName
	m.ExecOrder = dto.ExecOrder
	return m
}
