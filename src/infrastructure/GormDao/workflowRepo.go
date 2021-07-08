package GormDao

import (
	"JStock/src/domain/models/repos"

	"gorm.io/gorm"
)

// NewIFlowTemplateRepo


type WorkFlowRepo struct {
	db *gorm.DB
}

func NewWorkFlowRepoo(db *gorm.DB) *WorkFlowRepo {
	return &WorkFlowRepo{db: db}
}
func (s *WorkFlowRepo) New(model repos.IModel) error {
	return s.db.Table("t_workflow_template").Create(model).Error
}