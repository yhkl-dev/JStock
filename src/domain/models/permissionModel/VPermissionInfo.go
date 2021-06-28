package permissionmodel

type VPermInfo struct {
	ID       int    `json:"id" gorm:"column:id"`
	PermName string `json:"perm_name" gorm:"column:perm_name"`
}

func NewVPermInfo() *VPermInfo {
	return &VPermInfo{}
}
