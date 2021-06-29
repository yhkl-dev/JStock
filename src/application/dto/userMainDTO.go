package dto

import "time"

type (
	UserMainRequest struct {
		ID int `uri:"id" binding:"required"`
	}

	UserAddRequest struct {
		UserID          string `json:"user_id"`
		UserPassword    string `json:"user_password"`
		UserNameZh      string `json:"user_name_zh"`
		UserNameEn      string `json:"user_name_en"`
		UserMobilePhone string `json:"user_mobile_phone" `
		UserTelePhone   string `json:"user_telephone"`
		UserEmail       string `json:"user_email"`
		Remark          string `json:"remark"`
	}

	UserUpdateRequest struct {
		ID              int    `json:"id"`
		UserPassword    string `json:"user_password"`
		UserNameZh      string `json:"user_name_zh"`
		UserNameEn      string `json:"user_name_en"`
		UserMobilePhone string `json:"user_mobile_phone" `
		UserTelePhone   string `json:"user_telephone"`
		UserEmail       string `json:"user_email"`
		Remark          string `json:"remark"`
	}

	UserListRequest struct {
		UserID     string `form:"user_id"`
		UserNameEn string `form:"user_name_en"`
		UserNameZh string `form:"user_name_zh"`
		Page       int    `form:"page"`
		PageSize   int    `form:"page_size"`
	}
)

type (
	UserMainResponse struct {
		ID              int           `json:"id"`
		UserID          string        `json:"user_id"`
		UserNameZh      string        `json:"user_name_zh"`
		UserNameEn      string        `json:"user_name_en"`
		UserMobilePhone string        `json:"user_mobile_phone"`
		UserTelePhone   string        `json:"user_telephone" `
		UserEmail       string        `json:"user_email" `
		Remark          string        `json:"remark"`
		RoleInfos       []interface{} `json:"role_infos"`
		CreateAt        string        `json:"create_at" `
	}
	UserListResponse []UserMainResponse
	UserAddResponse  struct {
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
)
