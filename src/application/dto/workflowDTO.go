package dto

type (
	WorkFlowAddRequest struct {
		WorkFlowTemplateID int `json:"workflow_template_id"`
		MNCTemplateID      int `json:"mnc_template_id"`
	}
)

type (
	WorkFlowAddResponse struct {
		ID                 int `json:"id"`
		WorkFlowTemplateID int `json:"workflow_template_id"`
		MNCTemplateID      int `json:"mnc_template_id"`
	}
)