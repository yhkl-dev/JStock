package userRoleMap

import "JStock/src/domain/models/repos"

type UserRoleMapAttrFunc func(model *UserRoleMap)
type UserRoleMapAttrFuncs []UserRoleMapAttrFunc

func (s UserRoleMapAttrFuncs) apply(u *UserRoleMap) {
	for _, f := range s {
		f(u)
	}
}

func WithUserID(id int) UserRoleMapAttrFunc {
	return func(model *UserRoleMap) {
		model.UserID = id
	}
}

func WithRoleID(id int) UserRoleMapAttrFunc {
	return func(model *UserRoleMap) {
		model.RoleID = id
	}
}

type UserRoleMap struct {
	ID       int                `json:"id" gorm:"column:id"`
	UserID   int                `json:"user_id" gorm:"column:user_id"`
	RoleID   int                `json:"role_id" gorm:"column:role_id"`
	RoleName string             `json:"role_name" gorm:"column:role_name"`
	Repo     repos.IUserRoleMap `gorm:"-" json:"-"`
}

func (s *UserRoleMap) List(userID int) (interface{}, error) {
	return s.Repo.List(userID)
}

func (s *UserRoleMap) NewMap(UserID, RoleID int) error {
	return s.Repo.NewMap(UserID, RoleID)
}

func NewVUserRoleMap() *UserRoleMap {
	return &UserRoleMap{}
}

func New(attrs ...UserRoleMapAttrFunc) *UserRoleMap {
	c := &UserRoleMap{}
	UserRoleMapAttrFuncs(attrs).apply(c)
	return c
}
