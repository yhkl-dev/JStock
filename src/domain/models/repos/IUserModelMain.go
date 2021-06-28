package repos

type IUserModelMain interface {
	FindByID(IModel) error
	New(IModel) error
	GetUsetList(userID, userNameZh, userNameEn string, page, pageSize int) ([]IModel, error)
}
