package configs

import (
	"JStock/src/application/assembler"
	"JStock/src/application/services"
)

type RoleMainServiceConfig struct{}

func NewRoleMainServiceConfig() *RoleMainServiceConfig {
	return &RoleMainServiceConfig{}
}

func (s *RoleMainServiceConfig) RoleMainService() *services.RoleMainService {
	return &services.RoleMainService{
		AssRoleMainReq: &assembler.RoleMainRequest{},
		AssRoleMainRsp: &assembler.RoleMainResponse{},
		AssRoleListReq: &assembler.RoleListRequest{},
		AssRoleListRsp: &assembler.RoleListResponse{},
		AssRoleAddReq:  &assembler.RoleAddRequest{},
		AssRoleAddRsp:  &assembler.RoleAddResponse{},
	}
}
