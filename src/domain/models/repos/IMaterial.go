package repos

type IMaterialMainModel interface {
	New(IModel) error
	Load(IModel) error
}