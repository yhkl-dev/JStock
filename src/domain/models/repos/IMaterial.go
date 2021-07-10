package repos

type IMaterialMainModel interface {
	New(IModel) error
	Load(IModel) error
	QueryMaterialList(m IModel, page int, pageSize int) (interface{}, error)
	QueryImportancyLevelList() (interface{}, error)
}

type IPlantModel interface {
	New(IModel) error
	Load(IModel) error
	QueryPlantList(m IModel, page int, pageSize int) (interface{}, error)
}

type IPlantTechModel interface {
	New(IModel) error
	Load(IModel) error
	QueryPlantTechList(m IModel, page int, pageSize int) (interface{}, error)
}

type IMaterialGroupModel interface {
	New(IModel) error
	Load(IModel) error
	QueryMaterialGroupList(m IModel, page int, pageSize int) (interface{}, error)
}
