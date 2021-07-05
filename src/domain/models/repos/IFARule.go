package repos

type IFARuleModel interface {
	New(IModel) error
	List(IModel) (interface{}, error)
}
