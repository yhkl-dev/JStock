package setter

import (
	"fmt"
	"jstock/src/data/mappers"
	"jstock/src/models/roleModel"
	"jstock/src/result"
)

var RoleSetter IRoleSetter

type IRoleSetter interface {
	SaveRole(role *roleModel.Role) *result.ErrorResult
	UpdateRole(role *roleModel.Role) *result.ErrorResult
	DeleteRole(id int) *result.ErrorResult
}

func init() {
	RoleSetter = NewRoleSetterImp()
}

type RoleSetterImp struct {
	RoleMapper *mappers.RoleMapper
}

func NewRoleSetterImp() *RoleSetterImp {
	return &RoleSetterImp{RoleMapper: &mappers.RoleMapper{}}
}

func (s *RoleSetterImp) SaveRole(role *roleModel.Role) *result.ErrorResult {
	addRoleObject := s.RoleMapper.AddRole(role)
	err := mappers.Mappers(addRoleObject).Exec(func() error {
		err := addRoleObject.Exec().Error
		if err != nil {
			return err
		}
		// do something else

		return nil
	})
	return result.Result(role, err)
}

func (s *RoleSetterImp) UpdateRole(role *roleModel.Role) *result.ErrorResult {
	ret := s.RoleMapper.UpdateRole(role).Exec()
	return result.Result(role, ret.Error)
}

func (s *RoleSetterImp) DeleteRole(id int) *result.ErrorResult {
	res := s.QueryRoleUsers(id)
	if res != 0 && res != -1 {
		return result.Result("", fmt.Errorf(fmt.Sprintf("can not delete role, role %d has been used", res)))
	}
	ret := s.RoleMapper.DeleteRoleByID(id).Exec()
	return result.Result(ret.RowsAffected, ret.Error)
}

func (s *RoleSetterImp) QueryRoleUsers(id int) int {
	x := &struct {
		Id int
	}{}
	s.RoleMapper.QueryRoleUserById(id).Query().First(x)
	return x.Id
}
