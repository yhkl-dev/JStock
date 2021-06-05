package setter

import (
	"jstock/src/data/mappers"
	"jstock/src/models/userModel"
	"jstock/src/result"
)

var UserSetter IUserSetter

type IUserSetter interface {
	SaveUser(user *userModel.User) *result.ErrorResult
	UpdateUser(user *userModel.User) *result.ErrorResult
	DeleteUser(id int) *result.ErrorResult
	DeleteUserRoleByID(id int) *result.ErrorResult
}

func init() {
	UserSetter = NewUserSetterImp()
}

type UserSetterImp struct {
	UserMapper *mappers.UserMapper
}

func NewUserSetterImp() *UserSetterImp {
	return &UserSetterImp{UserMapper: &mappers.UserMapper{}}
}

func (s *UserSetterImp) SaveUser(user *userModel.User) *result.ErrorResult {
	addUserObject := s.UserMapper.AddUser(user)
	err := mappers.Mappers(addUserObject).Exec(func() error {
		err := addUserObject.Exec().Error
		if err != nil {
			return err
		}
		// do something else

		return nil
	})
	return result.Result(user, err)
}

func (s *UserSetterImp) UpdateUser(user *userModel.User) *result.ErrorResult {
	ret := s.UserMapper.UpdateUser(user).Exec()
	return result.Result(user, ret.Error)
}

func (s *UserSetterImp) DeleteUserRoleByID(id int) *result.ErrorResult {
	ret := s.UserMapper.DeteleUserRole(id).Exec()
	return result.Result(ret.RowsAffected, ret.Error)
}

func (s *UserSetterImp) DeleteUser(id int) *result.ErrorResult {
	// ret := s.UserMapper.DeleteUserByID(id).Exec()
	deleteUserObject := s.UserMapper.DeleteUserByID(id)
	deleteUserRole := s.UserMapper.DeteleUserRole(id)
	err := mappers.Mappers(deleteUserObject, deleteUserRole).Exec(func() error {
		err := deleteUserObject.Exec().Error
		if err != nil {
			return err
		}
		err = deleteUserRole.Exec().Error
		if err != nil {
			return err
		}
		// do something else

		return nil
	})
	return result.Result("success", err)
}
