package dto

type (
	ItemAddRequest struct {
		TemplateID int    `json:"template_id"`
		ItemName   string `json:"item_name"`
		ExecOrder  int    `json:"exec_order"`
		RoleID     int    `json:"role_id"`
	}
	ItemUpdateRequest struct {
		ID         int    `json:"id"`
		TemplateID int    `json:"template_id"`
		ItemName   string `json:"item_name"`
		ExecOrder  int    `json:"exec_order"`
		RoleID     int    `json:"role_id"`
	}
)

type (
	ItemAddResponse struct {
		TemplateName string `json:"template_name"`
		ItemName     string `json:"item_name"`
		ExecOrder    int    `json:"exec_order"`
		RoleName     string `json:"role_name"`
	}
	ItemUpdateResponse struct {
		ID           int    `json:"id"`
		TemplateName string `json:"template_name"`
		ItemName     string `json:"item_name"`
		ExecOrder    int    `json:"exec_order"`
		RoleName     string `json:"role_name"`
	}
	ItemMainResponse struct {
		ID           int    `json:"id"`
		TemplateName string `json:"template_name"`
		ItemName     string `json:"item_name"`
		ExecOrder    int    `json:"exec_order"`
		RoleName     string `json:"role_name"`
	}
	ItemListResponse []ItemMainResponse
)
