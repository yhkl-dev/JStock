package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type MaterialGroupService struct {
	AssPlantMainReq *assembler.MaterialGroupListRequest
	AssPlantMainRsp *assembler.MaterialGroupListResponse
	DB              *gorm.DB `inject:"-"`
}

func (s *MaterialGroupService) ListMaterialGroup(req *dto.MaterialGroupListRequest) *dto.MaterialGroupListResponse {
	materialGroup := materialmodel.NewMaterialGroup()
	material := materialmodel.New()
	plant := materialmodel.NewPlantModel()
	plantTech := materialmodel.NewPlantTenchModel()

	repo1 := GormDao.NewMaterialRepo(s.DB)
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	repo3 := GormDao.NewMaterialGroupRepo(s.DB)
	repo4 := GormDao.NewPlantTechRepo(s.DB)

	f := frontmaterial.NewFrontMaterialAgg(material, repo1, plant, repo2, materialGroup, repo3, plantTech, repo4)
	results, err := f.QueryMaterialGroupList(materialGroup, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssPlantMainRsp.D2M_MaterialGroupList(results)
}
