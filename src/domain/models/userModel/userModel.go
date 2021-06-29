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
	CreateAt *VUserCreatAt        `json:"create_at" gorm:"embedded"`
	RoleInfo []interface{}        `gorm:"-"`
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

func (s *UserModelMain) NewUser(user interface{}) (int, error) {
	u := user.(*UserModelMain)
	err := s.Repo.New(u)
	return u.ID, err
}

func (s *UserModelMain) List(userID, userNameZh, userNameEn string, page int, pageSize int) (interface{}, error) {
	userList, err := s.Repo.UserList(userID, userNameZh, userNameEn, page, pageSize)
	return userList, err
}

func (s *UserModelMain) UpdateUser() error {
	return s.Repo.UpdateUser(s)
}

func (s *UserModelMain) SetRole(roleID ...int) error {
	return s.Repo.SetRole(roleID...)
}

func (s *UserModelMain) GetRole() (interface{}, error) {
	return s.Repo.GetRole(s)
}

func WithRepo(repo repos.IUserModelMain) UserModelMainAttrFunc {
	return func(model *UserModelMain) {
		model.Repo = repo
	}
}

func New(attrs ...UserModelMainAttrFunc) *UserModelMain {
	c := &UserModelMain{
		UserInfo: NewVUserInfo(),
		CreateAt: NewVUserCreatAt(),
	}
	UserModelMainAttrFuncs(attrs).apply(c)
	return c
}
