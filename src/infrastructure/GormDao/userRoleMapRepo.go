package GormDao

import (
	userRoleMap "JStock/src/domain/models/userRoleMapModel"

	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
)

type UserRoleMapRepo struct {
	db *gorm.DB
}

func NewUserRoleMapRepo(db *gorm.DB) *UserRoleMapRepo {
	return &UserRoleMapRepo{db: db}
}

func (s *UserRoleMapRepo) List(userID int) (interface{}, error) {
	var mapModel = make([]*userRoleMap.UserRoleMap, 0)
	sql, args, _ := squirrel.
		Select("turm.id, turm.user_id, turm.role_id, tr.role_name").
		From("t_user_role_mapping turm, t_roles tr").
		Where("tr.id = turm.role_id and turm.user_id = ?", userID).
		ToSql()
	err := s.db.Raw(sql, args...).Scan(&mapModel).Error
	if err != nil {
		return nil, err
	}
	return mapModel, nil
}

func (s *UserRoleMapRepo) NewMap(userID, roleID int) error {
	m := userRoleMap.New(
		userRoleMap.WithUserID(userID),
		userRoleMap.WithRoleID(roleID),
	)

	return s.db.Table("t_user_role_mapping").Create(m).Error
}
