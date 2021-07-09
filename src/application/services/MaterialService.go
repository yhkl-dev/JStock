package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type MaterialService struct {
	AssMaterialMainReq *assembler.MaterialMainRequest
	AssMaterialMainRsp *assembler.MaterialMainResponse
	AssMaterialListReq *assembler.MaterialListRequest
	AssMaterialListRsp *assembler.MaterialListResponse
	DB                 *gorm.DB `inject:"-"`
}

func (s *MaterialService) GetMaterialInfo() error {
	return nil
}

func (s *MaterialService) CreateMaterial(req *dto.MaterialAddRequest) *dto.MaterialMainResponse {
	m := s.AssMaterialMainReq.D2M_MaterialMain(req)
	repo := GormDao.NewMaterialRepo(s.DB)
	plant := materialmodel.NewPlantModel()
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	f := frontmaterial.NewFrontMaterialAgg(m, repo, plant, repo2)
	f.CreateMaterial(m)
	return s.AssMaterialMainRsp.D2M_MaterialMain(f)
}

func (s *MaterialService) ListMaterial(req *dto.MatertialListRequest) *dto.MaterialListResponse {
	material := s.AssMaterialListReq.D2M_MaterialMain(req)
	repo := GormDao.NewMaterialRepo(s.DB)
	plant := materialmodel.NewPlantModel()
	repo2 := GormDao.NewPlantMainRepo(s.DB)
	f := frontmaterial.NewFrontMaterialAgg(material, repo, plant, repo2)
	results, err := f.QueryMaterialList(material, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssMaterialListRsp.D2M_MaterialList(results)
}
