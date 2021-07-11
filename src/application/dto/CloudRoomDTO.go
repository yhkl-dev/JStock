package dto

type (
	CloudRoomMainRequest struct {
		CloudRoomID string `json:"cloud_root_id" `
		RoomName    string `json:"room_name" `
		Owner       string `json:"owner" `
	}

	CloudRoomListRequest struct {
		CloudRoomID string `json:"cloud_root_id" `
		RoomName    string `json:"room_name" `
		Owner       string `json:"owner" `
		Page        int    `json:"page"`
		PageSize    int    `json:"page_size"`
	}
)

type (
	CloudRoomMainResponse struct {
		ID          int    `json:"id" `
		CloudRoomID string `json:"cloud_root_id" `
		RoomName    string `json:"room_name" `
		Owner       string `json:"owner" `
		CreateAt    string `json:"create_at" `
	}

	CloudRoomListResponse []CloudRoomMainResponse
)
