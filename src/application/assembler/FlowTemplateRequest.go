package assembler

import (
	"JStock/src/application/dto"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
)

type FlowTemplateAddRequest struct{}

func (s *FlowTemplateAddRequest) D2M_FlowTemplateMain(dto *dto.FlowTemplateAddRequest) *workflowtemplate.WorkFlowTemplate {
	m := workflowtemplate.NewWorkFlowTemplate()
	m.FlowName = dto.FlowName
	m.FlowType = dto.FlowType
	return m
}

type FlowTemplateListRequest struct{}

func (s *FlowTemplateListRequest) D2M_FlowTemplateList(dto *dto.FlowTemplateListRequest) *workflowtemplate.WorkFlowTemplate {
	m := workflowtemplate.NewWorkFlowTemplate()
	m.FlowName = dto.FlowName
	m.FlowType = dto.FlowType
	return m

}
