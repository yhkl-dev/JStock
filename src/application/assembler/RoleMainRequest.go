package assembler

import (
	"JStock/src/application/dto"
	rolemodel "JStock/src/domain/models/roleModel"
	"time"
)

type RoleMainRequest struct{}

func (s *RoleMainRequest) D2M_RoleMain(dto *dto.RoleMainRequest) *rolemodel.RoleModelMain {
	m := rolemodel.New()
	m.ID = dto.ID
	return m
}

type RoleListRequest struct{}

func (s *RoleListRequest) D2M_RoleList(dto *dto.RoleListRequest) *rolemodel.RoleModelMain {
	m := rolemodel.New()
	m.ID = dto.ID
	m.RoleInfo.RoleName = dto.RoleName
	m.RoleInfo.Desciption = dto.Desciption
	m.RoleInfo.ParentRoleID = dto.ParentRoleID
	return m
}

type RoleAddRequest struct{}

func (s *RoleAddRequest) D2M_Role(dto *dto.RoleAddRequest) *rolemodel.RoleModelMain {
	m := rolemodel.New()
	m.RoleInfo.RoleName = dto.RoleName
	m.RoleInfo.Desciption = dto.Desciption
	m.RoleInfo.ParentRoleID = dto.ParentRoleID
	m.RoleInfo.CreateAt = time.Now()
	return m
}
