package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type PlantService struct {
	AssPlantMainReq     *assembler.PlantAddRequest
	AssPlantMainRsp     *assembler.PlantAddResponse
	AssPlantListReq     *assembler.PlantListRequest
	AssPlantListRsp     *assembler.PlantListResponse
	AssPlantCodeListReq *assembler.PlantTechCodeListRequest
	AssPlantCodeListRsp *assembler.PlantTechCodeListResponse
	DB                  *gorm.DB `inject:"-"`
}

func (s *PlantService) CreatePlant(req *dto.PlantAddRequest) *dto.PlantMainResponse {
	material := materialmodel.New()
	repo1 := GormDao.NewMaterialRepo(s.DB)
	p := s.AssPlantMainReq.D2M_PlantModel(req)
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	g := materialmodel.NewMaterialGroup()
	repo3 := GormDao.NewMaterialGroupRepo(s.DB)
	pt := materialmodel.NewPlantTenchModel()
	repo4 := GormDao.NewPlantTechRepo(s.DB)
	f := frontmaterial.NewFrontMaterialAgg(material, repo1, p, repo2, g, repo3, pt, repo4)
	f.CreatePlant(p)
	return s.AssPlantMainRsp.D2M_PlantMain(f)
}

func (s *PlantService) ListPlant(req *dto.PlantListRequest) *dto.PlantListResponse {
	material := materialmodel.New()
	plant := s.AssPlantListReq.D2M_PlantModel(req)
	materialGroup := materialmodel.NewMaterialGroup()
	plantTech := materialmodel.NewPlantTenchModel()

	repo1 := GormDao.NewMaterialRepo(s.DB)
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	repo3 := GormDao.NewMaterialGroupRepo(s.DB)
	repo4 := GormDao.NewPlantTechRepo(s.DB)

	f := frontmaterial.NewFrontMaterialAgg(material, repo1, plant, repo2, materialGroup, repo3, plantTech, repo4)
	results, err := f.QueryPlantList(plant, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssPlantListRsp.D2M_PlantList(results)
}

func (s *PlantService) ListPlantTechCode(req *dto.PlantTechCodeListRequest) *dto.PlantTechCodeListResponse {
	material := materialmodel.New()
	plant := materialmodel.NewPlantModel()
	plantTechCode := s.AssPlantCodeListReq.D2M_PlantTechCodeModel(req)
	materialGroup := materialmodel.NewMaterialGroup()

	repo1 := GormDao.NewMaterialRepo(s.DB)
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	repo3 := GormDao.NewMaterialGroupRepo(s.DB)
	repo4 := GormDao.NewPlantTechRepo(s.DB)

	f := frontmaterial.NewFrontMaterialAgg(material, repo1, plant, repo2, materialGroup, repo3, plantTechCode, repo4)
	results, err := f.QueryPlantTechList(plantTechCode, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssPlantCodeListRsp.D2M_PlantTechCodeList(results)
}
