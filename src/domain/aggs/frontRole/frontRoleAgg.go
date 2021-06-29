package frontrole

import (
	"JStock/src/domain/models/repos"
	rolemodel "JStock/src/domain/models/roleModel"
	"JStock/src/infrastructure/Errors"
)

type FrontRoleAgg struct {
	RoleMain     *rolemodel.RoleModelMain
	RoleMainRepo repos.IRoleModelMain
}

func NewFrontRoleAgg(roleMain *rolemodel.RoleModelMain, roleMainRepo repos.IRoleModelMain) *FrontRoleAgg {
	if roleMain == nil {
		panic("Error Root RoleMain")
	}
	fr := &FrontRoleAgg{RoleMain: roleMain, RoleMainRepo: roleMainRepo}
	fr.RoleMain.Repo = roleMainRepo
	return fr
}

func (s *FrontRoleAgg) QueryDetail() error {
	if s.RoleMain.ID <= 0 {
		return Errors.NewNotFoundIDError("RoleMain")
	}
	err := s.RoleMain.Load()
	if err != nil {
		return Errors.NewNotFoundDataError("RoleMain", err.Error())
	}
	return nil
}

func (s *FrontRoleAgg) QueryRoleList(roleMain *rolemodel.RoleModelMain, page int, pageSize int) (interface{}, error) {
	return s.RoleMain.List(roleMain.RoleInfo.RoleName, roleMain.RoleInfo.ParentRoleID, page, pageSize)
}

func (s *FrontRoleAgg) CreateRole(role *rolemodel.RoleModelMain) error {
	return s.RoleMain.NewRole(role)
}
