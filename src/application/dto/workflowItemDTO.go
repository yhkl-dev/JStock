package dto

type (
	ItemAddRequest struct {
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
)
