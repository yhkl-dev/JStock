package frontmaterial

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"
	"JStock/src/infrastructure/Errors"
)

type FrontMaterialAgg struct {
	MaterialMain     *materialmodel.MaterialModel
	MaterialMainRepo repos.IMaterialMainModel
}

func NewFrontMaterialAgg(material *materialmodel.MaterialModel, repo repos.IMaterialMainModel) *FrontMaterialAgg {
	if material == nil {
		panic("Error Root MaterialMain")
	}
	f := &FrontMaterialAgg{MaterialMain: material, MaterialMainRepo: repo}
	f.MaterialMain.Repo = repo
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
