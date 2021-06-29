package rolemodel

import "JStock/src/domain/models/repos"

type RoleModelMainAttrFunc func(model *RoleModelMain)
type RoleModelMainAttrFuncs []RoleModelMainAttrFunc

func (s RoleModelMainAttrFuncs) apply(u *RoleModelMain) {
	for _, f := range s {
		f(u)
	}
}

type RoleModelMain struct {
	ID       int                  `json:"id" gorm:"column:id"`
	RoleInfo *VRoleInfo           `jons:"role_info" gorm:"embedded"`
	Repo     repos.IRoleModelMain `gorm:"-"`
}

func (*RoleModelMain) Name() string {
	return "RoleModelMain"
}

func (s *RoleModelMain) New() error {
	return s.Repo.New(s)
}

func (s *RoleModelMain) Load() error {
	return s.Repo.FindByID(s)
}

func (s *RoleModelMain) List(roleName string, parentRoleID int, page, pageSize int) (interface{}, error) {
	return s.Repo.List(roleName, parentRoleID, page, pageSize)
}

func (s *RoleModelMain) NewRole(role interface{}) error {
	return s.Repo.NewRole(role.(*RoleModelMain))
}

func WithRepo(repo repos.IRoleModelMain) RoleModelMainAttrFunc {
	return func(model *RoleModelMain) {
		model.Repo = repo
	}
}

func New(attrs ...RoleModelMainAttrFunc) *RoleModelMain {
	c := &RoleModelMain{
		RoleInfo: NewVRoleInfo(),
	}
	RoleModelMainAttrFuncs(attrs).apply(c)
	return c
}
