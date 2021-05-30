package material

import (
	"jstock/apps/base"
	"jstock/apps/users"
)

// 负责人信息
type RespInfo struct {
	ID             int
	TechnologyCode base.TechnologyCode
	User           users.User
}

// Material
type Material struct {
	ID                         int
	MaterialNumber             string
	TechCodes                  []base.TechCode // 工厂信息
	HandOverType               string          // 是否是通用件 同车间通用/同工厂通用/跨工厂通用/unique 非通用件
	MaterialDescriptionEN      string
	MaterialDescriptionZH      string
	ManufacturerName           string
	ManufacturerPartNumber     string
	ManufacturerModel          string
	SupplierName               string
	MaterialGroup              string
	Unit                       string
	Calibration                string
	Repairable                 string
	Material                   string
	CCCorCCCRelated            string
	PositionNumber             string
	MaterialMainClassification string
	MaterialSubClassification  string
	ManufactureModelOld1       string
	ManufactureModelOld2       string
	ManufacturePNOld1          string
	ManufacturePNOld2          string
	Dimension                  string
	RespInfos                  []RespInfo // 负责人信息
	MaterialSpecialTreatment   string
	ImportancyLevel            base.ImportancyLevel // nsm abc
	SurplusPoint               string
	InstallQty                 string
	MPRemark                   string // material planner remark
	TechRemark                 string
}
