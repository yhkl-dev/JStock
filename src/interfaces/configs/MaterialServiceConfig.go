package configs

import (
	"JStock/src/application/assembler"
	"JStock/src/application/services"
)

type MaterialMainServiceConfig struct{}

func NewMaterialMainServiceConfig() *MaterialMainServiceConfig {
	return &MaterialMainServiceConfig{}
}

func (s *MaterialMainServiceConfig) MaterialMainService() *services.MaterialService {
	return &services.MaterialService{
		AssMaterialMainReq: &assembler.MaterialMainRequest{},
		AssMaterialMainRsp: &assembler.MaterialMainResponse{},
		AssMaterialListReq: &assembler.MaterialListRequest{},
		AssMaterialListRsp: &assembler.MaterialListResponse{},
	}
}
