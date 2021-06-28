package GormDao

import (
	"JStock/src/domain/models/repos"
	usermodel "JStock/src/domain/models/userModel"

	"gorm.io/gorm"
)

type UserMainRepo struct {
	db *gorm.DB
}

func NewUserMainRepo(db *gorm.DB) *UserMainRepo {
	return &UserMainRepo{db: db}
}
func (s *UserMainRepo) FindByID(model repos.IModel) error {
	return s.db.Table("t_users").Where("id = ? ", model.(*usermodel.UserModelMain).ID).First(model).Error
}

func (s *UserMainRepo) New(model repos.IModel) error {
	return s.db.Table("t_users").Create(model).Error
}

func (s *UserMainRepo) UserList(models []*repos.IModel, userID, userNameZh, userNameEn string, page, pageSize int) ([]usermodel.UserModelMain, error) {
	var userModels []usermodel.UserModelMain
	err := s.db.Table("t_users").
		Where("user_id = ? and user_name_en = ? and user_name_zh = ? ", userID, userNameEn, userNameZh).
		Offset(page).Limit(pageSize).
		Find(&userModels).Error
	if err != nil {
		return nil, err
	}
	return userModels, nil
}
