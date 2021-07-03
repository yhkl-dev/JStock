package frontuser

import (
	"JStock/src/domain/models/repos"
	usermodel "JStock/src/domain/models/userModel"
	userRoleMap "JStock/src/domain/models/userRoleMapModel"
	"JStock/src/infrastructure/Errors"
	"JStock/src/interfaces/utils"
	"errors"
)

type FrontUserAgg struct {
	UserMain        *usermodel.UserModelMain
	UserRoleMap     *userRoleMap.UserRoleMap
	UserMainRepo    repos.IUserModelMain
	UserRoleMapRepo repos.IUserRoleMap
}

func NewFrontUserAgg(userMain *usermodel.UserModelMain, userRoleMap *userRoleMap.UserRoleMap, userMainRepo repos.IUserModelMain, userRoleMapRepo repos.IUserRoleMap) *FrontUserAgg {
	if userMain == nil {
		panic("Error Root UserMain")
	}
	fu := &FrontUserAgg{UserMain: userMain, UserRoleMap: userRoleMap, UserMainRepo: userMainRepo, UserRoleMapRepo: userRoleMapRepo}
	fu.UserRoleMap.Repo = userRoleMapRepo
	fu.UserMain.Repo = userMainRepo
	return fu
}

func (s *FrontUserAgg) QueryDetail() error {
	if s.UserMain.ID <= 0 {
		return Errors.NewNotFoundIDError("UserMain")
	}
	err := s.UserMain.Load()
	if err != nil {
		return Errors.NewNotFoundDataError("UserMain", err.Error())
	}
	results, err := s.UserRoleMap.List(s.UserMain.ID)
	if err != nil {
		return Errors.NewNotFoundDataError("UserMain", err.Error())
	}
	s.UserMain.RoleInfo = results
	return nil
}

func (s *FrontUserAgg) QueryUserList(userMain *usermodel.UserModelMain, page int, pageSize int) (interface{}, error) {
	results, err := s.UserMain.List(userMain.UserID, userMain.UserInfo.UserNameZh, userMain.UserInfo.UserNameEn, page, pageSize)
	if err != nil {
		return nil, err
	}
	for _, m := range results.([]*usermodel.UserModelMain) {

		res, err := s.UserRoleMap.List(m.ID)
		if err != nil {
			return nil, Errors.NewNotFoundDataError("UserRoleMap", err.Error())
		}
		m.RoleInfo = res
	}
	return results, nil
}

func (s *FrontUserAgg) CreateUser(user *usermodel.UserModelMain) (int, error) {
	return s.UserMain.NewUser(user)
}

func (s *FrontUserAgg) CreateUserWithRole(user *usermodel.UserModelMain, roleID ...int) error {
	id, err := s.UserMain.NewUser(user)
	if id <= 0 || err != nil {
		return Errors.NewNotFoundIDError("UserMain")
	}
	for _, role := range roleID {
		err := s.UserRoleMap.NewMap(id, role)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FrontUserAgg) UpdateUser() error {
	if s.UserMain.ID <= 0 {
		return Errors.NewNotFoundIDError("UserMain")
	}
	err := s.UserMain.UpdateUser()
	if err != nil {
		return Errors.NewNotFoundDataError("UserMain", err.Error())
	}
	return nil
}

func (s *FrontUserAgg) LoginFunc(pass string) error {
	err := s.UserMain.LoadUserID()
	if err != nil {
		return Errors.NewNotFoundDataError("UserMain", err.Error())
	}
	if s.UserMain.UserInfo.UserPassword != utils.Md5(pass) {
		return errors.New("password error")
	}
	results, err := s.UserRoleMap.List(s.UserMain.ID)
	if err != nil {
		return Errors.NewNotFoundDataError("UserRoleMap", err.Error())
	}
	s.UserMain.RoleInfo = results
	return nil
}
