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


func (s *PlantMainRepo) QueryPlantList(model repos.IModel, page, pageSize int) (interface{}, error)  {
	return nil, nil 
}


