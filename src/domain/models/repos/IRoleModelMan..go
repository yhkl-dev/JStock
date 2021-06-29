package repos

type IRoleModelMain interface {
	FindByID(IModel) error
	New(IModel) error
	NewRole(role interface{}) error
	List(roleName string, parentRoleID, page, pageSize int) (interface{}, error)
}
