package materialPlanning

type VPlanningInfo struct {
	TemplateName           string  `json:"template_name" gorm:"column:template_name"`                     // 模板名称
	MaterialNumberID       int     `json:"material_number_id"  gorm:"column:material_number_id"`          // 零件号
	PlantID                int     `json:"plant" gorm:"column:column:plant"`                              // 工厂代码
	MaterialGroupID        int     `json:"material_group_id" gorm:"column:material_group_id"`             // 物料组
	EngineerID             int     `json:"engineer" gorm:"column:engineer"`                               // 维修计划负责人
	TechnologyCode         int     `json:"plant_tech_code_id" gorm:"column:plant_tech_code_id"`           // 车间代码
	PurchaseTypeID         int     `json:"purchase_type" gorm:"column:purchase_type"`                     // 采购类型
	ImportancyLevelID      int     `json:"importancy_level_id" gorm:"column:importancy_level_id"`         // 等级
	GICategoryID           int     `json:"gi_category_id" gorm:"column:gi_category_id"`                   // GI类别
	StorageLocation        string  `json:"storage_location" gorm:"column:storage_location"`               // 存储库位
	PartDescriptionZh      string  `json:"material_description_zh" gorm:"column:material_description_zh"` // 零件描述(中文）
	PartDescriptionEN      string  `json:"material_description_en" gorm:"column:material_description_en"` // 零件描述(英文)
	Dimension              string  `json:"dimension" gorm:"column:dimension"`                             // 技术数据
	MinStock               int     `json:"min_stock" gorm:"column:min_stock"`                             // 安全库存
	MaxStock               int     `json:"max_stock" gorm:"column:max_stock"`                             // 最大库存量
	HandoverStock          int     `json:"handover_stock" gorm:"column:handover_stock"`                   // 初始移交库存
	Unit                   string  `json:"unit" gorm:"column:unit"`                                       // 计量单位
	ManufacturerName       string  `json:"manufacture_name" gorm:"column:manufacture_name"`               // 制造商名称
	ManufacturerModel      string  `json:"manufacture_model" gorm:"column:manufacture_model"`             // 制造商的规格型号
	ManufacturerPartNumber string  `json:"manufacture_part_number" gorm:"column:manufacture_part_number"` // 制造商的订货号
	SupplierCode           string  `json:"supplier_code" gorm:"column:supplier_code"`                     // 供应商编码
	SupplierName           string  `json:"supplier_name" gorm:"column:supplier_name"`                     // 供应商名称
	SupplierModel          string  `json:"supplier_model" gorm:"column:supplier_model"`                   // 供应商的规格型号
	SupplierPartNumber     string  `json:"supplier_part_number" gorm:"column:supplier_part_number"`       // 供应商的订货号
	SupplierContactPerson  string  `json:"supplier_contact_person" gorm:"column:supplier_contact_person"` // 供应商联系人
	SupplierContactPhone   string  `json:"supplier_contact_phone" gorm:"column:supplier_contact_phone"`   // 供应商联系电话
	SupplierEmail          string  `json:"supplier_email" gorm:"column:supplier_email"`                   // 供应商邮箱
	FCNo                   string  `json:"fc_no" gorm:"column:fc_no"`                                     // 框架协议号码
	UnitPrice              float32 `json:"unit_price" gorm:"column:unit_price"`                           // 单价
	Currency               string  `json:"currency" gorm:"column:currency"`                               // 货币
	EquipmentName          string  `json:"equipment_name" gorm:"column:equipment_name"`                   // 设备名称
	Repairable             string  `json:"repairable" gorm:"column:repairable"`                           // 可修理的零件
	Calibration            string  `json:"calibration" gorm:"column:calibration"`                         // 校准件
	Drawing                string  `json:"drawing" gorm:"column:drawing"`                                 // 图纸号
	InstallationQuantity   string  `json:"installation_quantity" gorm:"column:installation_quantity"`     // 安装数量
	PartSpecialTreatment   string  `json:"part_special_treatment" gorm:"column:part_special_treatment"`   // 零件特殊处理
	SurplusStock           int     `json:"surplus_stock" gorm:"column:surplus_stock"`                     // 剩余库存
}

func NewVPlanningInfo() *VPlanningInfo {
	return &VPlanningInfo{}
}
