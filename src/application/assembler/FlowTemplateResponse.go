package assembler

import (
	"JStock/src/application/dto"
	frontworkflowtempate "JStock/src/domain/aggs/frontWorkflowTempate"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
)

type FlowTemplateAddResponse struct{}

func (s *FlowTemplateAddResponse) D2M_FlowTemplateInfo(uagg *frontworkflowtempate.FrontWorkFlowTemplateAgg) *dto.FlowTemplateAddResponse {
	m := &dto.FlowTemplateAddResponse{}
	m.FlowName = uagg.WorkFlowTempateMain.FlowName
	m.FlowType = uagg.WorkFlowTempateMain.FlowType
	return m
}

type FlowTemplateListResponse struct{}

func (s *FlowTemplateListResponse) D2M_FlowTemplateLIst(uagg interface{}) *dto.FlowTemplateListResponse {
	listZ := make(dto.FlowTemplateListResponse, 0)
	resList := uagg.([]*workflowtemplate.WorkFlowTemplate)
	for _, flow := range resList {
		info := &dto.FlowMainResponse{}
		info.ID = flow.ID
		info.FlowName = flow.FlowName
		info.FlowType = flow.FlowType
		info.FlowItems = s.D2M_FlowItemInfo(flow)
		listZ = append(listZ, *info)
	}
	return &listZ
}

func (s *FlowTemplateListResponse) D2M_FlowItemInfo(uagg *workflowtemplate.WorkFlowTemplate) dto.ItemListResponse {
	info := make(dto.ItemListResponse, 0)
	if uagg.FlowItems != nil {
		for _, x := range uagg.FlowItems.([]*workflowtemplate.WorkFlowItemTemplate) {
			mapInfo := &dto.ItemMainResponse{}
			mapInfo.ExecOrder = x.ExecOrder
			mapInfo.ItemName = x.ItemName
			mapInfo.RoleName = x.RoleName
			mapInfo.TemplateName = x.TemplateName
			info = append(info, *mapInfo)
		}
	}
	return info
}
