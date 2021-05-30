package models

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"username" binding:"required,min=4"`
	UserPassword string `json:"password"`
}

func (u *User) TableName() string {
	return "users"
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
