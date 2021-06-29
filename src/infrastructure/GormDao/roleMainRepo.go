package GormDao

import "gorm.io/gorm"

type RoleMainRepo struct {
	db *gorm.DB
}

func NewRoleMainRepo(db *gorm.DB) *RoleMainRepo {
	return &RoleMainRepo{db: db}
}
