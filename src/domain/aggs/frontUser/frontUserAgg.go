package frontuser

import (
	"JStock/src/domain/models/repos"
	usermodel "JStock/src/domain/models/userModel"
	"JStock/src/infrastructure/Errors"
)

type FrontUserAgg struct {
	UserMain *usermodel.UserModelMain
	// RoleMain     *rolemodel.RoleModelMain
	UserMainRepo repos.IUserModelMain
	// RoleMainRepo repos.IRoleModelMain
}

// func NewFrontUserAgg(userMain *usermodel.UserModelMain, roleMain *rolemodel.RoleModelMain, userMainRepo repos.IUserModelMain, roleMainRepo repos.IRoleModelMain) *FrontUserAgg {
// 	if userMain == nil {
// 		panic("Error Root UserMain")
// 	}
// 	fu := &FrontUserAgg{UserMain: userMain, RoleMain: roleMain, UserMainRepo: userMainRepo, RoleMainRepo: roleMainRepo}
// 	fu.RoleMain.Repo = roleMainRepo
// 	fu.UserMain.Repo = userMainRepo
// 	return fu
// }

func NewFrontUserAgg(userMain *usermodel.UserModelMain, userMainRepo repos.IUserModelMain) *FrontUserAgg {
	if userMain == nil {
		panic("Error Root UserMain")
	}
	fu := &FrontUserAgg{UserMain: userMain, UserMainRepo: userMainRepo}
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
	return nil
}

func (s *FrontUserAgg) QueryUserList(userMain *usermodel.UserModelMain, page int, pageSize int) (interface{}, error) {
	return s.UserMain.List(userMain.UserID, userMain.UserInfo.UserNameZh, userMain.UserInfo.UserNameEn, page, pageSize)
}

func (s *FrontUserAgg) CreateUser(user *usermodel.UserModelMain) (int, error) {
	return s.UserMain.NewUser(user)
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
