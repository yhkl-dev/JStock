package repos

type IModel interface {
	Name() string
	Load() error
	New() error
}
