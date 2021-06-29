package assembler

import (
	"JStock/src/application/dto"
	frontrole "JStock/src/domain/aggs/frontRole"
	rolemodel "JStock/src/domain/models/roleModel"
)

type RoleMainResponse struct{}

func (s *RoleMainResponse) D2M_RoleMainInfo(uagg *frontrole.FrontRoleAgg) *dto.RoleMainResponse {
	role := &dto.RoleMainResponse{}
	role.ID = uagg.RoleMain.ID
	role.RoleName = uagg.RoleMain.RoleInfo.RoleName
	role.ParentRoleID = uagg.RoleMain.RoleInfo.ParentRoleID
	role.CreateAt = uagg.RoleMain.RoleInfo.CreateAt.Format("2006-01-02 15:04:05")
	return role
}

type RoleListResponse []RoleMainResponse

func (s *RoleListResponse) D2M_RoleListInfo(uagg interface{}) *dto.RoleListResponse {
	roleListZ := make(dto.RoleListResponse, 0)
	roleList := uagg.([]*rolemodel.RoleModelMain)
	for _, role := range roleList {
		x := &dto.RoleMainResponse{}
		x.ID = role.ID
		x.RoleName = role.RoleInfo.RoleName
		x.ParentRoleID = role.RoleInfo.ParentRoleID
		x.Desciption = role.RoleInfo.Desciption
		x.CreateAt = role.RoleInfo.CreateAt.Format("2006-01-02 15:04:05")
		roleListZ = append(roleListZ, *x)
	}
	return &roleListZ
}

type RoleAddResponse struct {
}

func (s *RoleAddResponse) D2M_AddRoleInfo(uagg *frontrole.FrontRoleAgg) *dto.RoleAddResponse {
	role := &dto.RoleAddResponse{}
	role.RoleName = uagg.RoleMain.RoleInfo.RoleName
	role.ParentRoleID = uagg.RoleMain.RoleInfo.ParentRoleID
	role.Desciption = uagg.RoleMain.RoleInfo.Desciption
	return role
}
