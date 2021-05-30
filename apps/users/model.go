package users

type User struct {
	ID       int
	UserName string
	RoleInfo Role
}

type Role struct {
	ID       int
	RoleName string //
	PermInfo Permission
}

type Permission struct {
	ID             int
	PermissionName string
}
