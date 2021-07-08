package materialmodel

import "time"

type VMaterialInfo struct {
	HandoverType               string    `json:"handover_type" gorm:"column:handover_type"` // 是否是通用件 同车间通用/同工厂通用/跨工厂通用/unique 非通用件
	MaterialDescriptionEN      string    `json:"material_description_en" gorm:"column:material_description_en"`
	MaterialDescriptionZH      string    `json:"material_description_zh" gorm:"column:material_description_zh"`
	ManufacturerName           string    `json:"manufacturer_name" gorm:"column:manufacturer_name"`
	ManufacturerPartNumber     string    `json:"manufacturer_part_number" gorm:"column:manufacturer_part_number"`
	ManufacturerModel          string    `json:"manufacturer_model" gorm:"column:manufacturer_model"`
	Unit                       string    `json:"unit" gorm:"column:unit"`
	Calibration                string    `json:"calibration" gorm:"column:calibration"`
	Repairable                 string    `json:"repairable" gorm:"column:repairable"`
	Material                   string    `json:"material" gorm:"column:material"`
	CCCorCCCRelated            string    `json:"ccc_or_ccc_related" gorm:"column:ccc_or_ccc_related"`
	PositionNumber             string    `json:"position_number" gorm:"column:position_number"`
	MaterialMainClassification string    `json:"material_main_classification" gorm:"column:material_main_classification"`
	MaterialSubClassification  string    `json:"material_sub_classification" gorm:"column:material_sub_classification"`
	ManufactureModelOld1       string    `json:"manufacture_model_old1" gorm:"column:manufacture_model_old1"`
	ManufactureModelOld2       string    `json:"manufacture_model_old2" gorm:"column:manufacture_model_old2"`
	ManufacturePNOld1          string    `json:"manufacture_pn_old1" gorm:"column:manufacture_pn_old1"`
	ManufacturePNOld2          string    `json:"manufacture_pn_old2" gorm:"column:manufacture_pn_old2"`
	Dimension                  string    `json:"dimension" gorm:"column:dimension"`
	MaterialSpecialTreatment   string    `json:"material_special_treatment" gorm:"column:material_special_treatment"`
	MPRemark                   string    `json:"mpr_remark" gorm:"column:mpr_remark" ` // material planner remark
	TechRemark                 string    `json:"tech_remark" gorm:"column:tech_remark"`
	SupplierName               string    `json:"supplier_name" gorm:"column:supplier_name"` // 不同工厂之间的供应商可能不同
	SurplusPoint               string    `json:"surplus_point" gorm:"column:surplus_point"`
	InstallQty                 string    `json:"install_qty" gorm:"column:install_qty"`
	CreateAt                   time.Time `json:"create_at" gorm:"column:creat_at"`
	ImportancyLevelID          int       `json:"importancy_level_id" gorm:"column:importancy_level_id"`
	ImportancyLevelName        string    `json:"importancy_level_name" gorm:"column:importancy_level_name"`
	MaterialGroupID            int       `json:"material_group_id" gorm:"column:material_group_id"`
	MaterialGroupName          string    `json:"material_group_name" gorm:"column:material_group_name"`
	RespInfo                   []int     `json:"resp_info" gorm:"-"`
	PlantID                    int       `json:"plant_id" gorm:"column:plant_id"`
	PlatnName                  string    `json:"plant_name" gorm:"column:plant_name"`
	PlantTechCodeID            []int     `json:"plant_tech_code_id" gorm:"column:plant_tech_code"`
}

func NewVMaterialInfo() *VMaterialInfo {
	return &VMaterialInfo{}
}
