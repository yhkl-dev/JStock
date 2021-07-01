package repos

type IUserRoleMap interface {
	// New(IModel) error
	// Update(IModel) error
	// Delete(IModel) error
	NewMap(UserID, RoleID int) error
	List(UserID int) (interface{}, error)
}