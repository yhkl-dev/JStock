package getter

import (
	"jstock/src/data/mappers"
	purchasetypemodel "jstock/src/models/purchaseTypeModel"
)

var PurchaseTypeGetter IPurchaseTypeGetter

type IPurchaseTypeGetter interface {
	GetPurchaseTypeList() (types []*purchasetypemodel.PurchaseType)
}

func init() {
	PurchaseTypeGetter = NewIPurchaseTypeGetterImp()
}

type PurchaseTypeGetterImp struct {
	purchaseTypeMapper *mappers.PurchaseTypeMapper
}

func NewIPurchaseTypeGetterImp() *PurchaseTypeGetterImp {
	return &PurchaseTypeGetterImp{purchaseTypeMapper: &mappers.PurchaseTypeMapper{}}
}

func (s *PurchaseTypeGetterImp) GetPurchaseTypeList() (types []*purchasetypemodel.PurchaseType) {
	s.purchaseTypeMapper.GetPurchaseTypeList().Query().Find(&types)
	return
}
