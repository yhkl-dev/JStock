package assembler

import (
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
	usermodel "JStock/src/domain/models/userModel"
	userRoleMap "JStock/src/domain/models/userRoleMapModel"
)

type UserMainResponse struct{}

func (s *UserMainResponse) D2M_UserMainInfo(uagg *frontuser.FrontUserAgg) *dto.UserMainResponse {
	userInfo := &dto.UserMainResponse{}
	userInfo.ID = uagg.UserMain.ID
	userInfo.UserID = uagg.UserMain.UserID
	userInfo.UserNameEn = uagg.UserMain.UserInfo.UserNameEn
	userInfo.UserNameZh = uagg.UserMain.UserInfo.UserNameZh
	userInfo.UserMobilePhone = uagg.UserMain.UserInfo.UserMobilePhone
	userInfo.UserTelePhone = uagg.UserMain.UserInfo.UserTelePhone
	userInfo.UserEmail = uagg.UserMain.UserInfo.UserEmail
	userInfo.Remark = uagg.UserMain.UserInfo.Remark
	userInfo.CreateAt = uagg.UserMain.CreateAt.GetFormat()
	userInfo.RoleInfo = s.D2M_UserRoleMapInfo(uagg)
	return userInfo
}

func (s *UserMainResponse) D2M_UserRoleMapInfo(uagg *frontuser.FrontUserAgg) dto.UserRoleListMapResponse {
	info := make(dto.UserRoleListMapResponse, 0)
	for _, x := range uagg.UserMain.RoleInfo.([]*userRoleMap.UserRoleMap) {
		mapInfo := &dto.UserRoleMapResponse{}
		mapInfo.ID = x.ID
		mapInfo.RoleName = x.RoleName
		mapInfo.RoleID = x.RoleID
		info = append(info, *mapInfo)
	}
	return info
}

type UserListResponse []UserMainResponse

func (s *UserListResponse) D2M_UserListInfo(uagg interface{}) *dto.UserListResponse {
	userListZ := make(dto.UserListResponse, 0)
	userList := uagg.([]*usermodel.UserModelMain)
	for _, user := range userList {
		userInfo := &dto.UserMainResponse{}
		userInfo.ID = user.ID
		userInfo.UserID = user.UserID
		userInfo.UserNameEn = user.UserInfo.UserNameEn
		userInfo.UserNameZh = user.UserInfo.UserNameZh
		userInfo.UserMobilePhone = user.UserInfo.UserMobilePhone
		userInfo.UserTelePhone = user.UserInfo.UserTelePhone
		userInfo.UserEmail = user.UserInfo.UserEmail
		userInfo.Remark = user.UserInfo.Remark
		userInfo.CreateAt = user.CreateAt.GetFormat()
		userInfo.RoleInfo = s.D2M_UserRoleMapInfo(user)
		userListZ = append(userListZ, *userInfo)
	}
	return &userListZ
}

func (s *UserListResponse) D2M_UserRoleMapInfo(uagg *usermodel.UserModelMain) dto.UserRoleListMapResponse {
	info := make(dto.UserRoleListMapResponse, 0)
	if uagg.RoleInfo != nil {
		for _, x := range uagg.RoleInfo.([]*userRoleMap.UserRoleMap) {
			mapInfo := &dto.UserRoleMapResponse{}
			mapInfo.ID = x.ID
			mapInfo.RoleName = x.RoleName
			mapInfo.RoleID = x.RoleID
			info = append(info, *mapInfo)
		}
	}
	return info
}

type UserAddResponse struct{}

func (s *UserAddResponse) D2M_AddUserInfo(uagg *frontuser.FrontUserAgg) *dto.UserAddResponse {
	u := &dto.UserAddResponse{}
	u.ID = uagg.UserMain.ID
	u.UserID = uagg.UserMain.UserID
	u.UserNameEn = uagg.UserMain.UserInfo.UserNameEn
	u.UserNameZh = uagg.UserMain.UserInfo.UserNameZh
	u.UserMobilePhone = uagg.UserMain.UserInfo.UserMobilePhone
	u.UserTelePhone = uagg.UserMain.UserInfo.UserTelePhone
	u.UserEmail = uagg.UserMain.UserInfo.UserEmail
	u.Remark = uagg.UserMain.UserInfo.Remark
	u.CreateAt = uagg.UserMain.CreateAt.CreateAt
	return u
}

type UserUpdateResponse struct{}

func (s *UserUpdateResponse) D2M_UpdatedUserInfo(uagg *frontuser.FrontUserAgg) *dto.UserUpdateResponse {
	u := &dto.UserUpdateResponse{}
	u.ID = uagg.UserMain.ID
	u.UserID = uagg.UserMain.UserID
	u.UserNameEn = uagg.UserMain.UserInfo.UserNameEn
	u.UserNameZh = uagg.UserMain.UserInfo.UserNameZh
	u.UserMobilePhone = uagg.UserMain.UserInfo.UserMobilePhone
	u.UserTelePhone = uagg.UserMain.UserInfo.UserTelePhone
	u.UserEmail = uagg.UserMain.UserInfo.UserEmail
	u.Remark = uagg.UserMain.UserInfo.Remark
	u.CreateAt = uagg.UserMain.CreateAt.CreateAt
	u.RoleInfo = uagg.UserMain.RoleInfo
	return u
}

type UserLoginResponse struct{}

func (s *UserLoginResponse) D2M_UserInfo(uagg *frontuser.FrontUserAgg, token string, err error) *dto.UserLoginResponse {
	u := &dto.UserLoginResponse{}
	if err != nil {
		u.Error = err.Error()
		return u
	}
	u.Token = token
	u.UserInfo.ID = uagg.UserMain.ID
	u.UserInfo.UserID = uagg.UserMain.UserID
	u.UserInfo.UserNameEn = uagg.UserMain.UserInfo.UserNameEn
	u.UserInfo.UserNameZh = uagg.UserMain.UserInfo.UserNameZh
	u.UserInfo.UserMobilePhone = uagg.UserMain.UserInfo.UserMobilePhone
	u.UserInfo.UserTelePhone = uagg.UserMain.UserInfo.UserTelePhone
	u.UserInfo.UserEmail = uagg.UserMain.UserInfo.UserEmail
	u.UserInfo.Remark = uagg.UserMain.UserInfo.Remark
	u.UserInfo.CreateAt = uagg.UserMain.CreateAt.CreateAt.Format("2006-01-02 15:04:05")
	u.UserInfo.RoleInfo = uagg.UserMain.RoleInfo
	u.Error = err.Error()
	return u
}
