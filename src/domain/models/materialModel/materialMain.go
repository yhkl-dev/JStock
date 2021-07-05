package materialmodel

type MaterialModelAttrFunc func(model *MaterialModel)
type MaterialModelAttrFuncs []MaterialModelAttrFunc

func (s MaterialModelAttrFuncs) apply(u *MaterialModel) {
	for _, f := range s {
		f(u)
	}
}

type MaterialModel struct {
	ID             int            `json:"id" gorm:"column:id"`
	MaterialNumber string         `json:"material_number" gorm:"column:material_number"`
	MaterialInfo   *VMaterialInfo `json:"material_info" gorm:"embedded"`
}

func New(attrs ...MaterialModelAttrFunc) *MaterialModel {
	c := &MaterialModel{
		MaterialInfo: NewVMaterialInfo(),
	}
	return c
}
