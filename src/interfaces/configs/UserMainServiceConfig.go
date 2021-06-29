package configs

import (
	"JStock/src/application/assembler"
	"JStock/src/application/services"
)

type UserMainServiceConfig struct{}

func NewUserMainServiceConfig() *UserMainServiceConfig {
	return &UserMainServiceConfig{}
}

func (s *UserMainServiceConfig) UserMainService() *services.UserMainService {
	return &services.UserMainService{
		AssUserMainReq: &assembler.UserMainRequest{},
		AssUserMainRsp: &assembler.UserMainResponse{},
		AssUserListReq: &assembler.UserListRequest{},
		AssUserListRsp: &assembler.UserListResponse{},
		AssUserAddReq:  &assembler.UserAddRequest{},
		AssUserAddRsp:  &assembler.UserAddResponse{},
	}
}
