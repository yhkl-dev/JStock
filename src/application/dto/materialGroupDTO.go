package dto

type (
	MaterialGroupAddRequest struct {
		GroupName string `json:"group_name" `
		Comment   string `json:"comment" `
	}

	MaterialGroupListRequest struct {
		Page     int `form:"page"`
		PageSize int `form:"page_size"`
	}
)

type (
	MaterialGroupMainResponse struct {
		ID        int    `json:"id"`
		GroupName string `json:"group_name" `
		Comment   string `json:"comment" `
	}

	MaterialGroupListResponse []MaterialGroupMainResponse
)
