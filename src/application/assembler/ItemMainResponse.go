package assembler

import (
	"JStock/src/application/dto"
	frontworkflowtempate "JStock/src/domain/aggs/frontWorkflowTempate"
)

type ItemMainResponse struct{}

func (s *ItemMainResponse) D2M_ItemMainInfo(uagg *frontworkflowtempate.FrontWorkFlowItemTemplateAgg) *dto.ItemAddResponse {
	r := &dto.ItemAddResponse{}
	r.ExecOrder = uagg.WorkFlowItemTemplateMain.ExecOrder
	r.RoleName = uagg.WorkFlowItemTemplateMain.RoleName
	r.ItemName = uagg.WorkFlowItemTemplateMain.ItemName
	r.TemplateName = uagg.WorkFlowItemTemplateMain.TemplateName
	return r
}
