package GormDao

import (
	"JStock/src/domain/models/repos"
	usermodel "JStock/src/domain/models/userModel"
	userRoleMap "JStock/src/domain/models/userRoleMapModel"

	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
)

type UserMainRepo struct {
	db *gorm.DB
}

func NewUserMainRepo(db *gorm.DB) *UserMainRepo {
	return &UserMainRepo{db: db}
}

func (s *UserMainRepo) GetRole(model repos.IModel) (interface{}, error) {
	return nil, nil
}

func (s *UserMainRepo) SetRole(roleId ...int) error {
	return s.db.Table("t_user_role_mapping").Error
}

func (s *UserMainRepo) FindByID(model repos.IModel) error {
	err := s.db.Table("t_users").Where("id = ? ", model.(*usermodel.UserModelMain).ID).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *UserMainRepo) New(model repos.IModel) error {
	return s.db.Table("t_users").Create(model).Error
}

func (s *UserMainRepo) UpdateUser(model repos.IModel) error {
	m := model.(*usermodel.UserModelMain)
	updateData := map[string]interface{}{
		"user_name_en":      m.UserInfo.UserNameEn,
		"user_name_zh":      m.UserInfo.UserNameEn,
		"user_password":     m.UserInfo.UserPassword,
		"user_mobile_phone": m.UserInfo.UserMobilePhone,
		"user_telephone":    m.UserInfo.UserTelePhone,
		"user_email":        m.UserInfo.UserEmail,
		"remark":            m.UserInfo.Remark,
	}
	sql, args, _ := squirrel.Delete("t_user_role_mapping").Where("user_id = ?", m.ID).ToSql()
	err := s.db.Exec(sql, args...).Error
	if err != nil {
		return err
	}
	roleInfo := m.RoleInfo.([]int)
	for _, roleID := range roleInfo {

		sql, args, _ := squirrel.Insert("t_user_role_mapping").Columns("user_id", "role_id").Values(m.ID, roleID).ToSql()
		err := s.db.Exec(sql, args...).Error
		if err != nil {
			panic(err.Error())
			// return err
		}
	}
	sql, args, _ = squirrel.Update("t_users").SetMap(updateData).ToSql()
	return s.db.Exec(sql, args...).Error
}

func (s *UserMainRepo) NewUser(model interface{}) (int, error) {
	m := model.(*usermodel.UserModelMain)
	err := s.db.Table("t_users").Create(m).Error
	if err != nil {
		return 0, err
	}
	roleInfo := m.RoleInfo.([]int)
	for _, roleID := range roleInfo {
		r := userRoleMap.New(
			userRoleMap.WithRoleID(roleID),
			userRoleMap.WithUserID(m.ID),
		)
		err := s.db.Table("t_user_role_mapping").Create(r).Error
		if err != nil {
			return 0, err
		}
	}
	return m.ID, nil
}

func (s *UserMainRepo) UserList(userID, userNameZh, userNameEn string, page int, pageSize int) (interface{}, error) {
	var userModels = make([]*usermodel.UserModelMain, 0)
	s.db = s.db.Table("t_users")
	if userID != "" {
		s.db = s.db.Where("user_id = ?", userID)
	}
	if userNameEn != "" {
		s.db = s.db.Where("user_name_en = ?", userNameEn)
	}
	if userNameZh != "" {
		s.db = s.db.Where("user_name_zh = ?", userNameZh)
	}
	if pageSize == 0 {
		pageSize = 10
	}
	err := s.db.Find(&userModels).Offset(page).Limit(pageSize).Error
	if err != nil {
		return nil, err
	}
	return userModels, nil
}
