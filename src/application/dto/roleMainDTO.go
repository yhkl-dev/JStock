package dto

type (
	RoleMainRequest struct {
		ID int `uri:"id" binding:"required"`
	}

	RoleListRequest struct {
		ID           int    `form:"id"`
		RoleName     string `form:"role_name"`
		Desciption   string `form:"description"`
		ParentRoleID int    `form:"parent_role_id"`
		Page         int    `form:"page"`
		PageSize     int    `form:"page_size"`
	}

	RoleAddRequest struct {
		RoleName     string `json:"role_name"`
		Desciption   string `json:"description"`
		ParentRoleID int    `json:"parent_role_id"`
	}

	RoleUpdateRequest struct {
		ID           int    `json:"id"`
		RoleName     string `json:"role_name"`
		Desciption   string `json:"desciption"`
		ParentRoleID int    `json:"parent_role_id"`
	}
)

type (
	RoleMainResponse struct {
		ID           int    `json:"id"`
		RoleName     string `json:"role_name"`
		Desciption   string `json:"desciption"`
		ParentRoleID int    `json:"parent_role_id"`
		CreateAt     string `json:"create_at"`
	}

	RoleListResponse []RoleMainResponse

	RoleAddResponse struct {
		RoleName     string `json:"role_name"`
		Desciption   string `json:"desciption"`
		ParentRoleID int    `json:"parent_role_id"`
	}
)
