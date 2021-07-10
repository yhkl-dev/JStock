package materialmodel

import (
	"JStock/src/domain/models/repos"
	"time"
)

type PlantModel struct {
	ID        int               `json:"id" gorm:"column:id"`
	PlantName string            `json:"plant_name" gorm:"column:plant_name"`
	PlantCode string            `json:"plant_code" gorm:"column:plant_code"`
	Comment   string            `json:"comment" gorm:"column:comment"`
	CreateAt  time.Time         `json:"create_at" gorm:"column:create_at"`
	Repo      repos.IPlantModel `gorm:"-"`
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

func (s *PlantModel) QueryPlantList(model repos.IModel, page, pageSize int) (interface{}, error) {
	return s.Repo.QueryPlantList(model, page, pageSize)
}

type PlantTenchModel struct {
	ID                int                   `json:"id" gorm:"column:id"`
	PlantTechCodeName string                `json:"plant_tech_code_name" gorm:"column:plant_tech_code_name"`
	PlantTechCode     string                `json:"plant_tech_code" gorm:"column:plant_tech_code"`
	Comment           string                `json:"comment" gorm:"column:comment"`
	PlantID           int                   `json:"plant_id" gorm:"column:plant_id"`
	CreateAt          time.Time             `json:"create_at" gorm:"column:create_at"`
	Repo              repos.IPlantTechModel `gorm:"-"`
}

func (*PlantTenchModel) Name() string {
	return "PlantTenchModel"
}

func NewPlantTenchModel() *PlantTenchModel {
	return &PlantTenchModel{}
}

func (s *PlantTenchModel) New() error {
	return s.Repo.New(s)
}

func (s *PlantTenchModel) Load() error {
	return s.Repo.Load(s)
}

func (s *PlantTenchModel) QueryPlantTechList(model repos.IModel, page, pageSize int) (interface{}, error) {
	return s.Repo.QueryPlantTechList(model, page, pageSize)
}
