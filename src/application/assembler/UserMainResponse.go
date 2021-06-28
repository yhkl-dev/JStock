package assembler

import (
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
)

type UserMainResponse struct{}

func (s *UserMainResponse) D2M_UserMainInfo(uagg *frontuser.FrontUserAgg) *dto.UserMainResponse {
	userInfo := &dto.UserMainResponse{}
	userInfo.ID = uagg.UserMain.ID
	userInfo.UserNameEn = uagg.UserMain.UserInfo.UserNameEn
	userInfo.UserNameZh = uagg.UserMain.UserInfo.UserNameZh
	userInfo.UserMobilePhone = uagg.UserMain.UserInfo.UserMobilePhone
	userInfo.UserTelePhone = uagg.UserMain.UserInfo.UserTelePhone
	userInfo.UserEmail = uagg.UserMain.UserInfo.UserEmail
	userInfo.Remark = uagg.UserMain.UserInfo.Remark
	userInfo.CreateAt = uagg.UserMain.UserInfo.CreateAt
	return userInfo
}

type UserListResponse []UserMainResponse

func (s *UserListResponse) D2M_UserListInfo(uagg *frontuser.FrontUserAgg) *dto.UserListResponse {
	userList := &dto.UserListResponse{}
	return userList
}
