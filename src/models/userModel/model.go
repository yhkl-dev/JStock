package userModel

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"username" binding:"username"`
	UserPassword string `json:"password"`
	RoleInfo     []Role `json:"roles" gorm:"-"`
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

type Role struct {
	ID       int          `json:"id"`
	RoleName string       `json:"rolename"`
	PermInfo []Permission `json:"permissions"`
}

type Permission struct {
	ID             int    `json:"id"`
	PermissionName string `json:"permname"`
	Comments       string `json:"comments"`
}
