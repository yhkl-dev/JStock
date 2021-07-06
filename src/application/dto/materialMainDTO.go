package dto

type (
	MaterialAddRequest struct {
		// MaterialNumber string `json:"material_number"`
		ImportancyLevelID          int    `json:"importancy_level_id" `
		MaterialGroupID            int    `json:"material_group_id" `
		RespInfo                   []int  `json:"resp_info" `
		PlantID                    int    `json:"plant_id" `
		PlantTechCodeID            []int  `json:"plant_tech_code_id" `
		HandoverType               string `json:"handover_type" `
		MaterialDescriptionEN      string `json:"material_description_en" `
		MaterialDescriptionZH      string `json:"material_description_zh" `
		ManufacturerName           string `json:"manufacturer_name" `
		ManufacturerPartNumber     string `json:"manufacturer_part_number" `
		ManufacturerModel          string `json:"manufacturer_model" `
		Unit                       string `json:"unit" `
		Calibration                string `json:"calibration" `
		Repairable                 string `json:"repairable" `
		Material                   string `json:"material" `
		CCCorCCCRelated            string `json:"ccc_or_ccc_related" `
		PositionNumber             string `json:"position_number" `
		MaterialMainClassification string `json:"material_main_classification" `
		MaterialSubClassification  string `json:"material_sub_classification" `
		ManufactureModelOld1       string `json:"manufacture_model_old1" `
		ManufactureModelOld2       string `json:"manufacture_model_old2" `
		ManufacturePNOld1          string `json:"manufacture_pn_old1" `
		ManufacturePNOld2          string `json:"manufacture_pn_old2" `
		Dimension                  string `json:"dimension" `
		MaterialSpecialTreatment   string `json:"material_special_treatment" `
		MPRemark                   string `json:"mpr_remark" `
		TechRemark                 string `json:"tech_remark" `
		SupplierName               string `json:"supplier_name" `
		SurplusPoint               string `json:"surplus_point" `
		InstallQty                 string `json:"install_qty" `
	}
)

type (
	MaterialMainResponse struct {
		ImportancyLevelID          int         `json:"importancy_level_id" `
		ImportancyLevelName        string      `json:"importancy_level_name"`
		MaterialGroupID            int         `json:"material_group_id" `
		MaterialGroupName          string      `json:"material_group_name"`
		RespInfo                   interface{} `json:"resp_info" `
		PlantID                    int         `json:"plant_id" `
		PlatnName                  string      `json:"plant_name"`
		PlantTechCodeInfo          interface{} `json:"plant_tech_code_id" `
		HandoverType               string      `json:"handover_type" `
		MaterialDescriptionEN      string      `json:"material_description_en" `
		MaterialDescriptionZH      string      `json:"material_description_zh" `
		ManufacturerName           string      `json:"manufacturer_name" `
		ManufacturerPartNumber     string      `json:"manufacturer_part_number" `
		ManufacturerModel          string      `json:"manufacturer_model" `
		Unit                       string      `json:"unit" `
		Calibration                string      `json:"calibration" `
		Repairable                 string      `json:"repairable" `
		Material                   string      `json:"material" `
		CCCorCCCRelated            string      `json:"ccc_or_ccc_related" `
		PositionNumber             string      `json:"position_number" `
		MaterialMainClassification string      `json:"material_main_classification" `
		MaterialSubClassification  string      `json:"material_sub_classification" `
		ManufactureModelOld1       string      `json:"manufacture_model_old1" `
		ManufactureModelOld2       string      `json:"manufacture_model_old2" `
		ManufacturePNOld1          string      `json:"manufacture_pn_old1" `
		ManufacturePNOld2          string      `json:"manufacture_pn_old2" `
		Dimension                  string      `json:"dimension" `
		MaterialSpecialTreatment   string      `json:"material_special_treatment" `
		MPRemark                   string      `json:"mpr_remark" `
		TechRemark                 string      `json:"tech_remark" `
		SupplierName               string      `json:"supplier_name" `
		SurplusPoint               string      `json:"surplus_point" `
		InstallQty                 string      `json:"install_qty" `
	}
)