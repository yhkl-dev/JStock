package workflowtemplate

import "JStock/src/domain/models/repos"

type WorkFlowTemplateAttrFunc func(model *WorkFlowTemplate)
type WorkFlowTemplateAttrFuncs []WorkFlowTemplateAttrFunc

func (s WorkFlowTemplateAttrFuncs) apply(u *WorkFlowTemplate) {
	for _, f := range s {
		f(u)
	}
}

type WorkFlowTemplate struct {
	ID        int                          `json:"id" gorm:"column:id"`
	FlowName  string                       `json:"flow_name" grom:"column:flow_name"`
	FlowType  int                          `json:"flow_type" gorm:"column:flow_type"`
	FlowItems interface{}                  `json:"flow_items" gorm:"-" `
	Repo      repos.IWorkFlowTemplateModel `gorm:"-"`
}

func NewWorkFlowTemplate(attrs ...WorkFlowTemplateAttrFunc) *WorkFlowTemplate {
	c := &WorkFlowTemplate{}
	WorkFlowTemplateAttrFuncs(attrs).apply(c)
	return c
}

func (*WorkFlowTemplate) Name() string {
	return "WorkFlowTemplate"
}

func (s *WorkFlowTemplate) New() error {
	return s.Repo.New(s)
}

func (s *WorkFlowTemplate) Load() error {
	return s.Repo.FindByID(s)
}

func (s *WorkFlowTemplate) NewWorkFlowTemplte(model repos.IModel) error {
	return s.Repo.NewWorkFlowTemplte(model)
}

func (s *WorkFlowTemplate) ListTemplate(flowName string, flowtype, page, pageSize int) (interface{}, error) {
	return s.Repo.ListTemplate(flowName, flowtype, page, pageSize)
}

type WorkFlowItemTempateAttrFunc func(model *WorkFlowItemTemplate)
type WorkFlowItemTempateAttrFuncs []WorkFlowItemTempateAttrFunc

func (s WorkFlowItemTempateAttrFuncs) apply(u *WorkFlowItemTemplate) {
	for _, f := range s {
		f(u)
	}
}

func NewWorkFlowItemTempate(attrs ...WorkFlowItemTempateAttrFunc) *WorkFlowItemTemplate {
	c := &WorkFlowItemTemplate{}
	WorkFlowItemTempateAttrFuncs(attrs).apply(c)
	return c
}

type WorkFlowItemTemplate struct {
	ID           int                              `json:"id" gorm:"column:id"`
	TemplateID   int                              `json:"tempalte_id" gorm:"template_id"`
	TemplateName string                           `json:"template_name" gorm:"column:template_name"`
	ItemName     string                           `json:"item_name" gorm:"column:item_name"`
	ExecOrder    int                              `json:"exec_order" gorm:"column:exex_order"`
	RoleID       int                              `json:"role_id" gorm:"column:role_id"`
	RoleName     string                           `json:"role_name" gorm:"column:role_name"`
	Repo         repos.IWorkFlowItemTemplateModel `gorm:"-"`
	// NextRoleID   int                              `json:"next_role_id" gorm:"column:next_role_id"`
	// NextRoleName string                           `json:"next_role_name" gorm:"column:next_role_name"`
}

func (*WorkFlowItemTemplate) Name() string {
	return "WorkFlowItemTemplate"
}

func (s *WorkFlowItemTemplate) New() error {
	return s.Repo.New(s)
}

func (s *WorkFlowItemTemplate) Load() error {
	return s.Repo.FindByID(s)
}

func (s *WorkFlowItemTemplate) List(id int) (interface{}, error) {
	return s.Repo.List(id)
}
