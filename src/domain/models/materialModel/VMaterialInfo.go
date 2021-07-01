package materialmodel

type VMaterialInfo struct {
	HandoverType               string `json:"handover_type" gorm:"column:handover_type"` // 是否是通用件 同车间通用/同工厂通用/跨工厂通用/unique 非通用件
	MaterialDescriptionEN      string `json:"material_description_en"`
	MaterialDescriptionZH      string `json:"material_description_zh"`
	ManufacturerName           string `json:"manufacturer_name"`
	ManufacturerPartNumber     string `json:"manufacturer_part_number"`
	ManufacturerModel          string `json:"manufacturer_model"`
	SupplierName               string `json:"supplier_name"`
	Unit                       string `json:"unit"`
	Calibration                string `json:"calibration"`
	Repairable                 string `json:"repairable"`
	Material                   string `json:"material"`
	CCCorCCCRelated            string `json:"ccc_or_ccc_related" gorm:"column:ccc_or_ccc_related"`
	PositionNumber             string `json:"position_number"`
	MaterialMainClassification string `json:"material_main_classification"`
	MaterialSubClassification  string `json:"material_sub_classification"`
	ManufactureModelOld1       string `json:"manufacture_model_old1"`
	ManufactureModelOld2       string `json:"manufacture_model_old2"`
	ManufacturePNOld1          string `json:"manufacture_pn_old1"`
	ManufacturePNOld2          string `json:"manufacture_pn_old2"`
	Dimension                  string `json:"dimension"`
	MaterialSpecialTreatment   string `json:"material_special_treatment"`
	SurplusPoint               string `json:"surplus_point"`
	InstallQty                 string `json:"install_qty"`
	MPRemark                   string `json:"mpr_remark" gorm:"column:mpr_remark" ` // material planner remark
	TechRemark                 string `json:"tech_remark"`

	ImportancyLevel int   `json:"importancy_level_id" gorm:"column:importancy_level_id"`
	TechCodes       []int `json:"tech_codes" sql:"-"`
	RespInfo        []int `json:"resp_info" sql:"-"`
	MaterialGroup   int   `json:"material_group"`
}
