package GormDao

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"

	"gorm.io/gorm"
)

type PlantTechRepo struct {
	db *gorm.DB
}

func NewPlantTechRepo(db *gorm.DB) *PlantTechRepo {
	return &PlantTechRepo{db: db}
}

func (s *PlantTechRepo) Load(model repos.IModel) error {
	return s.db.Table("t_plant_tech_code").Where("id = ? ", model.(*materialmodel.PlantTenchModel).ID).First(model).Error
}

func (s *PlantTechRepo) New(model repos.IModel) error {
	return s.db.Table("t_plant_tech_code").Create(model).Error
}

func (s *PlantTechRepo) QueryPlantTechList(model repos.IModel, page, pageSize int) (interface{}, error) {
	m := model.(*materialmodel.PlantTenchModel)
	var res = make([]*materialmodel.PlantTenchModel, 0)
	db := s.db.Table("t_plant_tech_code")
	if m.PlantTechCode != "" {
		db = db.Where("plant_tech_code = ? ", m.PlantTechCode)
	}
	if m.PlantID != 0 {
		db = db.Where("plant_id = ? ", m.PlantID)
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
