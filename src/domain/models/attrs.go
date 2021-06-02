package models

type UserAttrFunc func(u *User)
type UserAttrFuncs []UserAttrFunc

func (s UserAttrFuncs) Apply(u *User) {
	for _, f := range s {
		f(u)
	}
}

func WithUserID(id int) UserAttrFunc {
	return func(u *User) {
		u.ID = id
	}
}

func WithUserName(username string) UserAttrFunc {
	return func(u *User) {
		u.UserName = username
	}
}

func WithUserPassword(password string) UserAttrFunc {
	return func(u *User) {
		u.UserName = password
	}
}