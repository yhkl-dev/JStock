package services

import "gorm.io/gorm"

type UserRoleMapService struct {
	DB *gorm.DB `inject:"-"`
}