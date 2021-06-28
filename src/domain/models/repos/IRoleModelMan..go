package repos

type IRoleModelMain interface {
	FindByID(IModel) error
	New(IModel) error
}
