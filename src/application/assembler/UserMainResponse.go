package assembler

import (
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
	usermodel "JStock/src/domain/models/userModel"
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

func (s *UserListResponse) D2M_UserListInfo(uagg interface{}) *dto.UserListResponse {
	userListZ := make(dto.UserListResponse, 0)
	userList := uagg.([]*usermodel.UserModelMain)
	for _, user := range userList {
		userInfo := &dto.UserMainResponse{}
		userInfo.ID = user.ID
		userInfo.UserNameEn = user.UserInfo.UserNameEn
		userInfo.UserNameZh = user.UserInfo.UserNameZh
		userInfo.UserMobilePhone = user.UserInfo.UserMobilePhone
		userInfo.UserTelePhone = user.UserInfo.UserTelePhone
		userInfo.UserEmail = user.UserInfo.UserEmail
		userInfo.Remark = user.UserInfo.Remark
		userInfo.CreateAt = user.UserInfo.CreateAt
		userListZ = append(userListZ, *userInfo)
	}
	return &userListZ
}
