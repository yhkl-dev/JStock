package assembler

import (
	"JStock/src/application/dto"
	"JStock/src/domain/models/workflow"
)

type WorkFlowAddRequest struct{}


func (s *WorkFlowAddRequest) D2M_WorkFlow(dto *dto.WorkFlowAddRequest) *workflow.WorkFlow {
	m := workflow.NewWorkFlow()
	m.WorkFlowINfo.MNCTemplateID = dto.MNCTemplateID
	m.WorkFlowINfo.WorkFlowTemplateID = dto.WorkFlowTemplateID
	return m
}