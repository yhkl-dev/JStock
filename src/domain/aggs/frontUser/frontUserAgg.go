package frontuser

import (
	"JStock/src/domain/models/repos"
	usermodel "JStock/src/domain/models/userModel"
	"JStock/src/infrastructure/Errors"
	"fmt"
)

type FrontUserAgg struct {
	UserMain     *usermodel.UserModelMain
	UserMainRepo repos.IUserModelMain
	// RoleMain     *rolemodel.RoleModelMain

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
	// err = s.RoleMain.Load()
	// if err != nil {
	// 	return Errors.NewNotFoundDataError("RoleMain", err.Error())
	// }
	return nil
}

func (s *FrontUserAgg) QueryUserList(userMain *usermodel.UserModelMain) (interface{}, error) {
	// return s.UserMain.Userlist(userID, userNameZh, userNameEn, page, pageSize)
	res, err := s.UserMain.List(userMain.UserID, userMain.UserInfo.UserNameZh, userMain.UserInfo.UserNameEn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("???", res)
	return s.UserMain.List(userMain.UserID, userMain.UserInfo.UserNameZh, userMain.UserInfo.UserNameEn)
}
