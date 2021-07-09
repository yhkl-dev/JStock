package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/infrastructure/GormDao"
	"fmt"

	"gorm.io/gorm"
)

type PlantService struct {
	AssPlantMainReq *assembler.PlantAddRequest
	AssPlantMainRsp *assembler.PlantAddResponse
	AssPlantListReq *assembler.PlantListRequest
	AssPlantListRsp *assembler.PlantListResponse
	DB              *gorm.DB `inject:"-"`
}

func (s *PlantService) CreatePlant(req *dto.PlantAddRequest) *dto.PlantMainResponse {
	material := materialmodel.New()
	repo1 := GormDao.NewMaterialRepo(s.DB)
	p := s.AssPlantMainReq.D2M_PlantModel(req)
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	g := materialmodel.NewMaterialGroup()
	repo3 := GormDao.NewMaterialGroupRepo(s.DB)
	f := frontmaterial.NewFrontMaterialAgg(material, repo1, p, repo2, g, repo3)
	f.CreatePlant(p)
	return s.AssPlantMainRsp.D2M_PlantMain(f)
}

func (s *PlantService) ListPlant(req *dto.PlantListRequest) *dto.PlantListResponse {
	material := materialmodel.New()
	fmt.Println(req.PlantCode, req.PlantName)
	plant := s.AssPlantListReq.D2M_PlantModel(req)

	repo1 := GormDao.NewMaterialRepo(s.DB)
	repo2 := GormDao.NewPlantMainRepo(s.DB)

	f := frontmaterial.NewFrontMaterialAgg(material, repo1, plant, repo2)
	results, err := f.QueryPlantList(plant, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssPlantListRsp.D2M_PlantList(results)
}
