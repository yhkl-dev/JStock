package GormDao

import (
	"JStock/src/domain/models/repos"

	"gorm.io/gorm"
)

type MaterilGroupRepo struct {
	db *gorm.DB
}

func NewMaterilGroupRepo(db *gorm.DB) *MaterilGroupRepo {
	return &MaterilGroupRepo{db: db}
}

func (s *MaterilGroupRepo) New(model repos.IModel) error {
	return s.db.Table("t_material_group").Create(model).Error
}
