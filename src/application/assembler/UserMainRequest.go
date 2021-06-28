package assembler

import (
	"JStock/src/application/dto"
	usermodel "JStock/src/domain/models/userModel"
)

type UserMainRequest struct{}

func (s *UserMainRequest) D2M_UserMain(dto *dto.UserMainRequest) *usermodel.UserModelMain {
	m := usermodel.New()
	m.ID = dto.ID
	return m
}

type UserListRequest struct{}

func (s *UserListRequest) D2M_UserList(dto *dto.UserListRequest) *usermodel.UserModelMain {
	m := usermodel.New()
	// res, _ := x.GetUsetList(dto.UserID, dto.UserNameZh, dto.UserNameEn, dto.Page, dto.PageSize)
	m.UserID = dto.UserID
	m.UserInfo.UserNameEn = dto.UserNameEn
	m.UserInfo.UserNameZh = dto.UserNameZh
	return m
}
