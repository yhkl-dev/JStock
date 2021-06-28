package repos

type IPermModelMain interface {
	FindByID(IModel) error
	New(IModel) error
}
