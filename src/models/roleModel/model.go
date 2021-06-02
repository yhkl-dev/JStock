package roleModel

import "time"

type Role struct {
	ID          int       `json:"id"`
	RoleName    string    `json:"role_name"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

func New(attr ...RoleAttrFunc) *Role {
	u := &Role{}
	RoleAttrFuncs(attr).Apply(u)
	return u
}

func (s *Role) TableName() string {
	return "t_roles"
}

func (s *Role) Mutate(attr ...RoleAttrFunc) *Role {
	RoleAttrFuncs(attr).Apply(s)
	return s
}
