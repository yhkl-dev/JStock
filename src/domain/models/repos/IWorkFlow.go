package repos

type IWorkFlow interface {
	New(IModel) error
}

type IWorkFlowModel interface {
	New(IModel) error
}

type IWorkFlowTemplateModel interface {
	New(IModel) error
	FindByID(IModel) error
	NewWorkFlowTemplte(IModel) error
	Update(IModel) error
	ListTemplate(flowName string, flowtype, page, pageSize int) (interface{}, error)
}

type IWorkFlowItemTemplateModel interface {
	New(IModel) error
	List(id int) (interface{}, error)
	FindByID(IModel) error
	Update(IModel) error
}
