package GormDao

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"

	"gorm.io/gorm"
)

type PlantMainRepo struct {
	db *gorm.DB
}

func NewPlantMainRepo(db *gorm.DB) *PlantMainRepo {
	return &PlantMainRepo{db: db}
}

func (s *PlantMainRepo) Load(model repos.IModel) error {
	return s.db.Table("t_plant").Where("id = ? ", model.(*materialmodel.PlantModel).ID).First(model).Error
}

func (s *PlantMainRepo) New(model repos.IModel) error {
	return s.db.Table("t_plant").Create(model).Error
}

func (s *PlantMainRepo) QueryPlantList(model repos.IModel, page, pageSize int) (interface{}, error) {
	m := model.(*materialmodel.PlantModel)
	var res = make([]*materialmodel.PlantModel, 0)
	db := s.db.Table("t_plant")
	if m.PlantCode != "" {
		db = db.Where("plant_code = ?", m.PlantCode)
	}
	if m.PlantName != "" {
		db = db.Where("plant_name = ?", m.PlantName)
	}
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
