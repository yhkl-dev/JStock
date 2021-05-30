package base

import "fmt"

// PurchaseType 1：OTB 2.PRIE 3.CallOFF
type PurchaseType struct {
	ID           int
	PurchaseName string // 1：OTB 2.PRIE 3.CallOFF
	Comments     string
}

// MaterialGroup 1: Spare part S ; 2.comsumale C ; 3.tool T ; 4.others O
type MaterialGroup struct {
	ID        int
	GroupName string
	Comments  string
	ShortName string
	Code      string // A B C D E...
}

// CurrencyType 货币种类
type CurrencyType struct {
	ID           int
	CurrencyName string
	Comments     string
}

// TechCode 工厂代码 p1 p2 p3
type TechCode struct {
	ID       int
	Name     string
	Comments string
}

// TechnologyCode 车间代码 1: at1; 2 at2; 3: at3; 4:bt1; 5: bt2; 6: bt3
type TechnologyCode struct {
	ID       int
	CodeName string
	Comments string
	TechCode TechCode
}

// ImportancyLevel 重要度 a b c
type ImportancyLevel struct {
	ID        int
	LevelName string
	Comments  string
}

// GICategory goods issue category 出库类型
/*
	1：strategy part 战略件;
	2: wear part 磨损件
	3：consumable part 消耗件
	4：overhawl part 大修件
	5. preventive maintenaince part 预防性维修件
*/
type GICategory struct {
	ID           int
	CategoryName string
	Comments     string
}

func (s *PurchaseType) String() string {
	return fmt.Sprintf("<%s>", s.PurchaseName)
}
