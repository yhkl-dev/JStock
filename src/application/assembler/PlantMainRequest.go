package assembler

import (
	"JStock/src/application/dto"
	materialmodel "JStock/src/domain/models/materialModel"
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
	m := materialmodel.NewPlantModel()
	m.PlantName = dto.PlantName
	m.PlantCode = dto.PlantCode
	return m
}

type PlantTechCodeListRequest struct{}

func (s *PlantTechCodeListRequest) D2M_PlantTechCodeModel(dto *dto.PlantTechCodeListRequest) *materialmodel.PlantTenchModel {
	m := materialmodel.NewPlantTenchModel()
	m.PlantTechCode = dto.PlantTechCode
	m.PlantID = dto.PlantID
	return m
}
