package repos

type IUserModelMain interface {
	FindByID(IModel) error
	New(IModel) error
	NewUser(user interface{}) (int, error)
	UpdateUser(IModel) error
	UserList(userID, userNameZh, userNameEn string, page, pageSize int) (interface{}, error)
	SetRole(roleID ...int) error
	GetRole(IModel) (interface{}, error)
}
