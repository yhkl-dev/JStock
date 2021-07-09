package configs

import (
	"JStock/src/application/assembler"
	"JStock/src/application/services"
)

type PlantMainServiceConfig struct{}

func NewPlantMainServiceConfig() *PlantMainServiceConfig {
	return &PlantMainServiceConfig{}
}

func (s *PlantMainServiceConfig) PlantMainService() *services.PlantService {
	return &services.PlantService{
		AssPlantMainReq: &assembler.PlantAddRequest{},
		AssPlantMainRsp: &assembler.PlantAddResponse{},
		AssPlantListReq: &assembler.PlantListRequest{},
		AssPlantListRsp: &assembler.PlantListResponse{},
	}
}
