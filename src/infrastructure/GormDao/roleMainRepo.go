package GormDao

import (
	"JStock/src/domain/models/repos"
	rolemodel "JStock/src/domain/models/roleModel"

	"gorm.io/gorm"
)

type RoleMainRepo struct {
	db *gorm.DB
}

func NewRoleMainRepo(db *gorm.DB) *RoleMainRepo {
	return &RoleMainRepo{db: db}
}

func (s *RoleMainRepo) FindByID(model repos.IModel) error {
	return s.db.Table("t_roles").Where("id = ? ", model.(*rolemodel.RoleModelMain).ID).First(model).Error
}

func (s *RoleMainRepo) New(model repos.IModel) error {
	return s.db.Table("t_roles").Create(model).Error
}

func (s *RoleMainRepo) NewRole(model interface{}) error {
	return s.db.Table("t_roles").Create(model.(*rolemodel.RoleModelMain)).Error
}

func (s *RoleMainRepo) List(roleName string, parentRoleID int, page int, pageSize int) (interface{}, error) {
	var roleModels = make([]*rolemodel.RoleModelMain, 0)
	s.db = s.db.Table("t_roles")
	if roleName != "" {
		s.db = s.db.Where("role_name = ?", roleName)
	}
	if parentRoleID != 0 {
		s.db = s.db.Where("parent_role_id = ?", parentRoleID)
	}
	if pageSize == 0 {
		pageSize = 10
	}
	err := s.db.Find(&roleModels).Offset(page).Limit(pageSize).Error
	if err != nil {
		return nil, err
	}
	return roleModels, nil
}
