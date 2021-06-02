package userModel

import "jstock/src/models/roleModel"

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"username" binding:"username"`
	UserPassword string `json:"password"`
	RoleName     string `json:"role_name"`
	Roles        []*roleModel.Role `json:"roles" gorm:"-"`
}

func (u *User) TableName() string {
	return "t_users"
}

func New(attr ...UserAttrFunc) *User {
	u := &User{}
	UserAttrFuncs(attr).Apply(u)
	return u
}

func (s *User) Mutate(attr ...UserAttrFunc) *User {
	UserAttrFuncs(attr).Apply(s)
	return s
}
