package dto

import "time"

type (
	UserMainRequest struct {
		ID int `uri:"id" binding:"required"`
		// UserID int `uri:"user_id" binding:"required"`
	}

	UserListRequest struct {
		UserID     string `json:"user_id"`
		UserNameEn string `json:"user_name_en"`
		UserNameZh string `json:"user_name_zh"`
		Page       int    `json:"page"`
		PageSize   int    `json:"page_size"`
	}
)

type (
	UserMainResponse struct {
		ID              int       `json:"id"`
		UserID          string    `json:"user_id"`
		UserNameZh      string    `json:"user_name_zh"`
		UserNameEn      string    `json:"user_name_en"`
		UserMobilePhone string    `json:"user_mobile_phone"`
		UserTelePhone   string    `json:"user_telephone" `
		UserEmail       string    `json:"user_email" `
		Remark          string    `json:"remark"`
		CreateAt        time.Time `json:"create_at" `
	}
	UserListResponse []UserMainResponse
)
