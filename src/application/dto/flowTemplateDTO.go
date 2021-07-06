package dto

type (
	FlowTemplateAddRequest struct {
		FlowName string `json:"flow_name"`
		FlowType int    `json:"flow_type"`
	}
	FlowTemplateUpdateRequest struct {
		ID       int    `json:"id"`
		FlowName string `json:"flow_name"`
		FlowType int    `json:"flow_type"`
	}

	FlowTemplateListRequest struct {
		FlowName string `form:"flow_name"`
		FlowType int    `form:"flow_type"`
		Page     int    `form:"page"`
		PageSize int    `form:"page_size"`
	}
)

type (
	FlowMainResponse struct {
		ID        int         `json:"id"`
		FlowName  string      `json:"flow_name"`
		FlowType  int         `json:"flow_type"`
		FlowItems interface{} `json:"flow_items"`
	}
	FlowTemplateAddResponse struct {
		FlowName string `json:"flow_name"`
		FlowType int    `json:"flow_type"`
	}

	FlowTemplateListResponse []FlowMainResponse
)
