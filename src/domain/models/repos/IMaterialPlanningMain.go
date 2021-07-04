package repos

type IMaterialPlanningModelMain interface {
	New(IModel) error
	Update(IModel) error
	PlanningList(userID, userNameZh, userNameEn string, page, pageSize int) (interface{}, error)
}
