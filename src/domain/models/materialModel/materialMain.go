package materialmodel

type MaterialModel struct {
	ID             int    `json:"id" gorm:"column:id"`
	MaterialNumber string `json:"material_number" gorm:"column:material_number"`
}
