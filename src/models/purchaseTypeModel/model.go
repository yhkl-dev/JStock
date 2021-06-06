package purchasetypemodel

type PurchaseType struct {
	ID           int    `json:"id"`
	PurchaseType string `json:"purchase_type"` // 1：OTB 2.PRIE 3.CallOFF
	Comments     string `json:"comments"`
}

func (u *PurchaseType) TableName() string {
	return "t_purchase_type"
}
