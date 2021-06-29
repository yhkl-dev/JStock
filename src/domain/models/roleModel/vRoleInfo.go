package rolemodel

import "time"

func WithRoleName(rolename string) RoleModelMainAttrFunc {
	return func(model *RoleModelMain) {
		model.RoleInfo.RoleName = rolename
	}
}

func WithParentRoleID(id int) RoleModelMainAttrFunc {
	return func(model *RoleModelMain) {
		model.RoleInfo.ParentRoleID = id
	}
}

func WithDesciption(desciption string) RoleModelMainAttrFunc {
	return func(model *RoleModelMain) {
		model.RoleInfo.Desciption = desciption
	}
}

type VRoleInfo struct {
	RoleName     string    `json:"role_name" gorm:"column:role_name"`
	Desciption   string    `json:"description" gorm:"column:description"`
	ParentRoleID int       `json:"parent_role_id" gorm:"column:parent_role_id"`
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at"`
}

func NewVRoleInfo() *VRoleInfo {
	return &VRoleInfo{}
}
