package usermodel

import "JStock/src/domain/models/repos"

type UserModelMainAttrFunc func(model *UserModelMain)
type UserModelMainAttrFuncs []UserModelMainAttrFunc

func (s UserModelMainAttrFuncs) apply(u *UserModelMain) {
	for _, f := range s {
		f(u)
	}
}

type UserModelMain struct {
	ID       int                  `json:"id" gorm:"column:id"`
	UserID   string               `json:"user_id" gorm:"user_id"`
	UserInfo *VUserInfo           `json:"user_info" gorm:"embedded"`
	Repo     repos.IUserModelMain `gorm:"-"`
}

func (*UserModelMain) Name() string {
	return "userModelMain"
}

func (s *UserModelMain) New() error {
	return s.Repo.New(s)
}

func (s *UserModelMain) Load() error {
	return s.Repo.FindByID(s)
}

func (s *UserModelMain) List(userID, userNameZh, userNameEn string) (interface{}, error) {
	userList, err := s.Repo.UserList(userID, userNameZh, userNameEn)
	return userList, err
}

func WithRepo(repo repos.IUserModelMain) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.Repo = repo
	}
}

func New(attrs ...UserModelMainAttrFunc) *UserModelMain {
	c := &UserModelMain{
		UserInfo: NewVUserInfo(),
	}
	UserModelMainAttrFuncs(attrs).apply(c)
	return c
}
