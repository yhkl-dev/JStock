package GormDao

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"

	"gorm.io/gorm"
)

type MaterilGroupRepo struct {
	db *gorm.DB
}

func NewMaterialGroupRepo(db *gorm.DB) *MaterilGroupRepo {
	return &MaterilGroupRepo{db: db}
}

func (s *MaterilGroupRepo) New(model repos.IModel) error {
	return s.db.Table("t_material_group").Create(model).Error
}

func (s *MaterilGroupRepo) Load(model repos.IModel) error {
	return s.db.Table("t_material_group").Where("id = ?", model.(*materialmodel.MaterialGroup).ID).First(model).Error
}

func (s *MaterilGroupRepo) QueryMaterialGroupList(model repos.IModel, page int, pageSize int) (interface{}, error) {
	var res = make([]*materialmodel.MaterialGroup, 0)
	db := s.db.Table("t_material_group")
	if pageSize == 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 0
	}
	err := db.Limit(pageSize).Offset(page).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
