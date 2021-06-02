package getter

import (
	"jstock/src/data/mappers"
	"jstock/src/models/roleModel"
)

var RoleGetter IRoleGetter

type IRoleGetter interface {
	GetRoleList() (roles []*roleModel.Role)
}

func init() {
	RoleGetter = NewIRoleGetterImp()
}

type RoleGetterImp struct {
	roleMappter *mappers.RoleMapper
}

func NewIRoleGetterImp() *RoleGetterImp {
	return &RoleGetterImp{roleMappter: &mappers.RoleMapper{}}
}

func (s *RoleGetterImp) GetRoleList() (roles []*roleModel.Role) {
	s.roleMappter.GetRoleList().Query().Find(&roles)
	return
}
