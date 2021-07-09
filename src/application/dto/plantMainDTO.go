package dto

type (
	PlantAddRequest struct {
		PlantName string `json:"plant_name" `
		PlantCode string `json:"plant_code" `
		Comment   string `json:"comment" `
	}

	PlantListRequest struct {
		PlantName string `form:"plant_name" `
		PlantCode string `form:"plant_code" `
		Page      int    `form:"page"`
		PageSize  int    `form:"page_size"`
	}
)

type (
	PlantMainResponse struct {
		ID        int    `json:"id" `
		PlantName string `json:"plant_name" `
		PlantCode string `json:"plant_code" `
		Comment   string `json:"comment" `
		CreateAt  string `json:"create_at" `
	}
	PlantListResponse []PlantMainResponse
)