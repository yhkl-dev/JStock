package roleModel

type RoleAttrFunc func(u *Role)
type RoleAttrFuncs []RoleAttrFunc

func (s RoleAttrFuncs) Apply(u *Role) {
	for _, f := range s {
		f(u)
	}
}

func WithRoleID(id int) RoleAttrFunc {
	return func(u *Role) {
		u.ID = id
	}
}

func WithUserName(rolename string) RoleAttrFunc {
	return func(u *Role) {
		u.RoleName = rolename
	}
}
