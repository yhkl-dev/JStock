package repos

type IMaterialMainModel interface {
	New(IModel) error
	Load(IModel) error
	QueryMaterialList(m IModel, page int, pageSize int) (interface{}, error)
}
