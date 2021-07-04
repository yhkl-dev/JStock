package materialPlanning

type VIndicator struct {
	MinRecommentation int    `json:"min_recommentation"`
	MaxRecommentation int    `json:"max_recommentation"`
	RiskLevelAlert    string `json:"risk_level"`
	HighValue         string `json:"high_value" gorm:"high_value"`
}

func NewVIndicator() *VIndicator {
	return &VIndicator{}
}
