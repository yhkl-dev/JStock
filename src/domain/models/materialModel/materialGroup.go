package materialmodel

import "JStock/src/domain/models/repos"

type MaterialGroup struct {
	ID        int                       `json:"id" gorm:"column:id"`
	GroupName string                    `json:"group_name" gorm:"column:group_name"`
	Comment   string                    `json:"comment" gorm:"column:comment"`
	Repo      repos.IMaterialGroupModel `gorm:"-"`
}

func NewMaterialGroup() *MaterialGroup {
	return &MaterialGroup{}
}

func (s *MaterialGroup) Name() string {
	return "MaterialGroup"
}

func (s *MaterialGroup) New() error {
	return s.Repo.New(s)
}

func (s *MaterialGroup) Load() error {
	return s.Repo.Load(s)
}

func (s *MaterialGroup) QueryMaterialGroupList(model repos.IModel, page int, pageSize int) (interface{}, error) {
	return s.Repo.QueryMaterialGroupList(model, page, pageSize)
}
