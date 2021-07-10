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
	PlantTechCodeListRequest struct {
		PlantID       int    `form:"plant_id"`
		PlantTechCode string `form:"plant_tech_code"`
		Page          int    `form:"page"`
		PageSize      int    `form:"page_size"`
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

	PlantTechCodeMainResponse struct {
		ID                int    `json:"id" `
		PlantID           int    `json:"plant_id" `
		PlantTechCodeName string `json:"plant_tech_code_name" `
		PlantTechCode     string `json:"plant_tech_code" `
		Comment           string `json:"comment" `
		CreateAt          string `json:"create_at" `
	}
	PlantTechCodeListResponse []PlantTechCodeMainResponse
)
