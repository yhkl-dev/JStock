package materialPlanning

import "JStock/src/domain/models/repos"

type MaterialPlanningModelMainAttrFunc func(model *MaterialPlanningModelMain)
type MaterialPlanningModelMainAttrFuncs []MaterialPlanningModelMainAttrFunc

func (s MaterialPlanningModelMainAttrFuncs) apply(u *MaterialPlanningModelMain) {
	for _, f := range s {
		f(u)
	}
}

const (
	Start           = 1
	Waiting         = 2
	MRHConfirming   = 3
	Pendding        = 4
	MRHConfirmed    = 5
	STRConfirming   = 6
	STRConfirmed    = 7
	OTBCollection   = 8
	PRProcessing    = 9
	OrderCreating   = 10
	OrderConfirming = 11
	OrderConfirmed  = 12
	Producing       = 13
	Delivery        = 14
	GoodsReady      = 15
	Invoice         = 16
)

type MaterialPlanningModelMain struct {
	ID           int                              `json:"id" gorm:"column:id"`
	Status       int                              `json:"status" gorm:"status"`
	Indicator    *VIndicator                      `json:"indicator" gorm:"embedded"`
	PlanningInfo *VPlanningInfo                   `json:"planning_info" gorm:"embedded"`
	Repo         repos.IMaterialPlanningModelMain `gorm:"-"`
}

func WithRepo(repo repos.IMaterialPlanningModelMain) MaterialPlanningModelMainAttrFunc {
	return func(model *MaterialPlanningModelMain) {
		model.Repo = repo
	}
}

func New(attrs ...MaterialPlanningModelMainAttrFunc) *MaterialPlanningModelMain {
	c := &MaterialPlanningModelMain{
		Indicator:    NewVIndicator(),
		PlanningInfo: NewVPlanningInfo(),
	}
	MaterialPlanningModelMainAttrFuncs(attrs).apply(c)
	return c
}
