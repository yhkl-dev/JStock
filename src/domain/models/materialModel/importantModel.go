package materialmodel

type ImportancyLevel struct {
	ID        int    `json:"id" gorm:"column:id"`
	LevelName string `json:"level_name" gorm:"column:level_name"`
	Comment   string `json:"comment" gorm:"column:comment"`
}

func NewImportancyLevel() *ImportancyLevel {
	return &ImportancyLevel{}

}

func (s *ImportancyLevel) Name() string {
	return "ImportancyLevel"
}

func (s *ImportancyLevel) Load() error {
	return nil
}

func (s *ImportancyLevel) New() error {
	return nil
}
