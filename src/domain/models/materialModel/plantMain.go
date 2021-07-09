package materialmodel

import (
	"JStock/src/domain/models/repos"
	"time"
)


type PlantModel struct {
	ID        int       `json:"id" gorm:"column:id"`
	PlantName string    `json:"plant_name" gorm:"column:plant_name"`
	PlantCode string    `json:"plant_code" gorm:"column:plant_code"`
	Comment   string    `json:"comment" gorm:"column:comment"`
	CreateAt  time.Time `json:"create_at" gorm:"column:create_at"`
	Repo repos.IPlantModel `gorm:"-"`
}

func (*PlantModel) Name() string {
	return "PlantModel"
}

func NewPlantModel() *PlantModel {
	return &PlantModel{}
}

func (s *PlantModel) New() error {
	return s.Repo.New(s)
}

func (s *PlantModel) Load() error {
	return s.Repo.Load(s)
}