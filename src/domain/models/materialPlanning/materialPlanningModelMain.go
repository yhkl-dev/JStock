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
	STRBalanced     = 8
	STRFinished     = 9
	OTBCollection   = 10
	PRProcessing    = 11
	OrderCreating   = 12
	OrderConfirming = 13
	OrderConfirmed  = 14
	Producing       = 15
	Delivery        = 16
	GoodsReady      = 17
	GoodsReceived   = 18
	Invoice         = 19
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
