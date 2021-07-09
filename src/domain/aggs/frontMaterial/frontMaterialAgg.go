package frontmaterial

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"
	"JStock/src/infrastructure/Errors"
)

type FrontMaterialAgg struct {
	MaterialMain     *materialmodel.MaterialModel
	PlantModelMain *materialmodel.PlantModel
	MaterialMainRepo repos.IMaterialMainModel
	PlantModelMainRepo repos.IPlantModel
}

func NewFrontMaterialAgg(material *materialmodel.MaterialModel, repo repos.IMaterialMainModel, plant *materialmodel.PlantModel, repo2 repos.IPlantModel) *FrontMaterialAgg {
	if material == nil {
		panic("Error Root MaterialMain")
	}
	f := &FrontMaterialAgg{MaterialMain: material, PlantModelMain: plant, MaterialMainRepo: repo, PlantModelMainRepo: repo2 }
	f.MaterialMain.Repo = repo
	f.PlantModelMain.Repo = repo2
	return f
}

func (s *FrontMaterialAgg) QueryDetail() error {
	if s.MaterialMain.ID <= 0 {
		return Errors.NewNotFoundIDError("MaterialMain")
	}
	err := s.MaterialMain.Load()
	if err != nil {
		return Errors.NewNotFoundDataError("MaterialMain", err.Error())
	}
	return nil
}

func (s *FrontMaterialAgg) CreateMaterial(m *materialmodel.MaterialModel) error {
	return s.MaterialMain.New()
}

func (s *FrontMaterialAgg) QueryMaterialList(m *materialmodel.MaterialModel, page, pageSize int) (interface{}, error) {
	return s.MaterialMain.QueryMaterialList(m, page, pageSize)
}

func (s *FrontMaterialAgg) CreatePlant(m *materialmodel.PlantModel) error {
	return s.PlantModelMain.New()
}

func (s FrontMaterialAgg) QueryPlantList(m *materialmodel.PlantModel, page, pageSize int) (interface{}, error) {
	return s.PlantModelMain.Repo.QueryPlantList(m, page, pageSize)
}