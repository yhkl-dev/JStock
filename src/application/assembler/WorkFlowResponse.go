package assembler

import (
	"JStock/src/application/dto"
	frontworkflow "JStock/src/domain/aggs/frontWorkflow"
)

type WorkFlowAddResponse struct{}


func (s *WorkFlowAddResponse) D2M_WorkFlow(agg *frontworkflow.FrontWorkFlowAgg) *dto.WorkFlowAddResponse {
	m := &dto.WorkFlowAddResponse{}
	m.ID = agg.WorkFlowMain.ID
	
	return m
}