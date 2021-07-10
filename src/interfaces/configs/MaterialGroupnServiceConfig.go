package configs

import (
	"JStock/src/application/assembler"
	"JStock/src/application/services"
)

type MaterialGroupnServiceConfig struct{}

func NewMaterialGroupnServiceConfig() *MaterialGroupnServiceConfig {
	return &MaterialGroupnServiceConfig{}
}

func (s *MaterialGroupnServiceConfig) MaterialGroupService() *services.MaterialGroupService {
	return &services.MaterialGroupService{
		AssPlantMainReq: &assembler.MaterialGroupListRequest{},
		AssPlantMainRsp: &assembler.MaterialGroupListResponse{},
	}
}
