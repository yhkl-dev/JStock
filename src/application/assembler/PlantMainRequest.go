package assembler

import (
	"JStock/src/application/dto"
	materialmodel "JStock/src/domain/models/materialModel"
	"fmt"
	"time"
)

type PlantAddRequest struct{}

func (s *PlantAddRequest) D2M_PlantModel(dto *dto.PlantAddRequest) *materialmodel.PlantModel {
	m := materialmodel.NewPlantModel()
	m.PlantName = dto.PlantName
	m.Comment = dto.Comment
	m.PlantCode = dto.PlantCode
	m.CreateAt = time.Now()
	return m
}

type PlantListRequest struct{}

func (s *PlantListRequest) D2M_PlantModel(dto *dto.PlantListRequest) *materialmodel.PlantModel {
	fmt.Println("xxxxxxxxx", dto.PlantCode)
	m := materialmodel.NewPlantModel()

	m.PlantName = dto.PlantName
	m.PlantCode = dto.PlantCode
	return m
}
