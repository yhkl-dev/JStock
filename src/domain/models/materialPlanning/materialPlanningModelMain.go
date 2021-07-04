package materialPlanning

import "JStock/src/domain/models/repos"

type MaterialPlanningModelMainAttrFunc func(model *MaterialPlanningModelMain)
type MaterialPlanningModelMainAttrFuncs []MaterialPlanningModelMainAttrFunc

func (s MaterialPlanningModelMainAttrFuncs) apply(u *MaterialPlanningModelMain) {
	for _, f := range s {
		f(u)
	}
}

type MaterialPlanningModelMain struct {
	ID           int                              `json:"id" gorm:"column:id"`
	Status       int                              `json:"status" gorm:"status"`
	Indicator    string                           `json:"indicator" gorm:"indicator"`
	PlanningInfo *VPlanningInfo                   `json:"planning_info" gorm:"embedded"`
	Repo         repos.IMaterialPlanningModelMain `gorm:"-"`
}

func WithRepo(repo repos.IMaterialPlanningModelMain) MaterialPlanningModelMainAttrFunc {
	return func(model *MaterialPlanningModelMain) {
		model.Repo = repo
	}
}

func New(attrs ...MaterialPlanningModelMainAttrFunc) *MaterialPlanningModelMain {
	c := &MaterialPlanningModelMain{}
	MaterialPlanningModelMainAttrFuncs(attrs).apply(c)
	return c
}
