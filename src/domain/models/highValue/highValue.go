package HighValue

type HighValueModelMain struct {
	ID            int     `json:"id" gorm:"column:id"`
	PlantID       int     `json:"plant_id" gorm:"column:plant_id"`
	PlantTechCode int     `json:"plant_tech_code_id" gorm:"column:plant_tech_code_id"`
	UnitPrice     float64 `json:"unit_price" gorm:"column:unit_price"`
	Currency      string  `json:"currency" gorm:"column:currency"`
}

func NewHighValueModelMainI() *HighValueModelMain {
	return &HighValueModelMain{}
}
