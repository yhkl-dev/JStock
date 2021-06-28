package repos

type IUserModelMain interface {
	FindByID(IModel) error
	New(IModel) error
	UserList(userID, userNameZh, userNameEn string) (interface{}, error)
}
