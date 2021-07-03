package dto

import "time"

type (
	UserMainRequest struct {
		ID int `uri:"id" binding:"required"`
	}

	UserLoginRequest struct {
		UserId   string `json:"user_id"`
		Password string `json:"password"`
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
		Roles           []int  `json:"role_info"`
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
		Roles           []int  `json:"role_info"`
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
	UserLoginResponse struct {
		Token    string           `json:"token"`
		UserInfo UserMainResponse `json:"user_info"`
		Error    string           `json:"error"`
	}

	UserMainResponse struct {
		ID              int         `json:"id"`
		UserID          string      `json:"user_id"`
		UserNameZh      string      `json:"user_name_zh"`
		UserNameEn      string      `json:"user_name_en"`
		UserMobilePhone string      `json:"user_mobile_phone"`
		UserTelePhone   string      `json:"user_telephone" `
		UserEmail       string      `json:"user_email" `
		Remark          string      `json:"remark"`
		RoleInfo        interface{} `json:"role_info"`
		CreateAt        string      `json:"create_at" `
	}

	UserAddResponse struct {
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

	UserUpdateResponse struct {
		ID              int         `json:"id"`
		UserID          string      `json:"user_id"`
		UserNameZh      string      `json:"user_name_zh"`
		UserNameEn      string      `json:"user_name_en"`
		UserMobilePhone string      `json:"user_mobile_phone"`
		UserTelePhone   string      `json:"user_telephone" `
		UserEmail       string      `json:"user_email" `
		Remark          string      `json:"remark"`
		RoleInfo        interface{} `json:"role_info"`
		CreateAt        time.Time   `json:"create_at" `
	}
	UserRoleMapResponse struct {
		ID       int    `json:"id"`
		RoleID   int    `json:"role_id"`
		RoleName string `json:"roleName"`
	}

	UserListResponse        []UserMainResponse
	UserRoleListMapResponse []UserRoleMapResponse
)
