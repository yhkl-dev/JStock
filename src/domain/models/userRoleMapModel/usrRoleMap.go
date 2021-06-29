package userRoleMap

type VUserRoleMap struct {
	ID     int `json:"id" gorm:"column:id"`
	UserID int `json:"user_id" gorm:"column:user_id"`
	RoleID int `json:"role_id" gorm:"column:role_id"`
}

func NewVUserROleMap() *VUserRoleMap {
	return &VUserRoleMap{}
}
