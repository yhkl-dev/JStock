package assembler

import (
	"JStock/src/application/dto"
	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
)


type PlantAddResponse struct {
}

func (s *PlantAddResponse) D2M_PlantMain(agg *frontmaterial.FrontMaterialAgg) *dto.PlantMainResponse {
	m := &dto.PlantMainResponse{}
	m.PlantName = agg.PlantModelMain.PlantName
	m.Comment= agg.PlantModelMain.Comment
	m.PlantCode = agg.PlantModelMain.PlantCode
	m.CreateAt = agg.PlantModelMain.CreateAt.Format("2006-01-02 15:04:05")
	return m
}

type PlantListResponse struct{}

func (s *PlantListResponse) D2M_PlantList(agg interface{}) *dto.PlantListResponse {
	plantList := make(dto.PlantListResponse, 0)
	resPlantList := agg.([]*materialmodel.PlantModel)
	for _, m := range resPlantList {
		info := &dto.PlantMainResponse{}
		info.ID = m.ID
		info.Comment = m.Comment
		info.PlantCode = m.PlantCode
		info.PlantName = m.PlantName
		info.CreateAt = m.CreateAt.Format("2006-01-02 15:04:05")
		plantList = append(plantList, *info)
	}
	return &plantList
}