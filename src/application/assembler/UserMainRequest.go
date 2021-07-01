package assembler

import (
	"JStock/src/application/dto"
	usermodel "JStock/src/domain/models/userModel"
	"JStock/src/infrastructure/common"
	"time"
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
	m.UserID = dto.UserID
	m.UserInfo.UserNameEn = dto.UserNameEn
	m.UserInfo.UserNameZh = dto.UserNameZh
	return m
}

type UserAddRequest struct{}

func (s *UserAddRequest) D2M_User(dto *dto.UserAddRequest) *usermodel.UserModelMain {
	m := usermodel.New()
	m.UserID = dto.UserID
	m.UserInfo.UserPassword = common.MD5Encrypt(dto.UserPassword)
	m.UserInfo.UserNameEn = dto.UserNameEn
	m.UserInfo.UserNameZh = dto.UserNameZh
	m.UserInfo.UserMobilePhone = dto.UserMobilePhone
	m.UserInfo.UserTelePhone = dto.UserTelePhone
	m.UserInfo.UserEmail = dto.UserEmail
	m.UserInfo.Remark = dto.Remark
	m.CreateAt.CreateAt = time.Now()
	m.RoleInfo = dto.Roles
	return m
}

type UserUpdateRequest struct{}

func (s *UserUpdateRequest) D2M_User(dto *dto.UserUpdateRequest) *usermodel.UserModelMain {
	m := usermodel.New()
	m.ID = dto.ID
	m.UserInfo.UserPassword = common.MD5Encrypt(dto.UserPassword)
	m.UserInfo.UserNameEn = dto.UserNameEn
	m.UserInfo.UserNameZh = dto.UserNameZh
	m.UserInfo.UserMobilePhone = dto.UserMobilePhone
	m.UserInfo.UserTelePhone = dto.UserTelePhone
	m.UserInfo.UserEmail = dto.UserEmail
	m.RoleInfo = dto.Roles
	return m
}
