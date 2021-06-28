package permissionmodel

import "JStock/src/domain/models/repos"

type PermissionModelAttrFunc func(model *PermissionModel)
type PermissionModelAttrFuncs []PermissionModelAttrFunc

func (s PermissionModelAttrFuncs) apply(u *PermissionModel) {
	for _, f := range s {
		f(u)
	}
}

type PermissionModel struct {
	ID   int                  `json:"id" gorm:"column:id"`
	Repo repos.IPermModelMain `gorm:"-"`
}

func (*PermissionModel) Name() string {
	return "PermissionModel"
}

func (s *PermissionModel) New() error {
	return s.Repo.New(s)
}

func (s *PermissionModel) Load() error {
	return s.Repo.FindByID(s)
}

func WithRepo(repo repos.IPermModelMain) PermissionModelAttrFunc {
	return func(model *PermissionModel) {
		model.Repo = repo
	}
}

func New(attrs ...PermissionModelAttrFunc) *PermissionModel {
	c := &PermissionModel{}
	PermissionModelAttrFuncs(attrs).apply(c)
	return c
}
