package materialmodel

import "JStock/src/domain/models/repos"

type MaterialModelAttrFunc func(model *MaterialModel)
type MaterialModelAttrFuncs []MaterialModelAttrFunc

func (s MaterialModelAttrFuncs) apply(u *MaterialModel) {
	for _, f := range s {
		f(u)
	}
}

type MaterialModel struct {
	ID             int                      `json:"id" gorm:"column:id"`
	MaterialNumber string                   `json:"material_number" gorm:"column:material_number"`
	MaterialInfo   *VMaterialInfo           `json:"material_info" gorm:"embedded"`
	Repo           repos.IMaterialMainModel `gorm:"-"`
}

func New(attrs ...MaterialModelAttrFunc) *MaterialModel {
	c := &MaterialModel{
		MaterialInfo: NewVMaterialInfo(),
	}
	MaterialModelAttrFuncs(attrs).apply(c)
	return c
}

func (s *MaterialModel) Load() error {
	return s.Repo.Load(s)
}

func (s *MaterialModel) New() error {
	return s.Repo.New(s)
}

func (*MaterialModel) Name() string {
	return "MaterialModel"
}