package handover

import (
	"jstock/apps/base"
	"jstock/apps/users"
)

type HandOver struct {
	ID int
}

// MNCTemplate Material Creatation
type MNCTemplate struct {
	ID                     int
	Plant                  string               // 工厂代码 *必填
	PurchaseType           base.PurchaseType    // 采购类型 *必填
	StorageLocation        string               // 存储库位
	PartNumber             string               // 零件号
	MaterialGroup          base.MaterialGroup   // 物料组 *必填
	PartDescriptionCN      string               // 零件描述(中文）*必填
	PartDescriptionEN      string               // 零件描述(英文) *必填
	Dimension              string               // 技术数据
	OrderPoint             int                  // 订货点 *必填
	OrderToPoint           int                  // 订到点 *必填
	HandoverStock          int                  // 初始移交库存 *必填
	Unit                   string               // 计量单位  *必填
	ManufacturerName       string               // 制造商名称 *必填
	ManufacturerModel      string               // 制造商的规格型号
	ManufacturerPartNumber string               // 制造商的订货号
	SupplierCode           string               // 供应商编码
	SupplierName           string               // 供应商名称 *必填
	SupplierModel          string               // 供应商的规格型号
	SupplierPartNumber     string               // 供应商的订货号
	SupplierContactPerson  string               // 供应商联系人 *必填
	SupplierContactPhone   string               // 供应商联系电话 *必填
	SupplierEmail          string               // 供应商邮箱 *必填
	FCNo                   string               // 框架协议号码
	UnitPrice              float32              // 单价 *必填
	Currency               base.CurrencyType    // 货币 *必填
	EquipmentName          string               // 设备名称 *必填
	Engineer               users.User           // 维修计划负责人 *必填
	TechnologyCode         base.TechnologyCode  // 车间代码 *必填
	Repairable             bool                 // 是否修理的零件 *必填 y/n
	Calibration            bool                 // 校准件 *必填 y/n y/n
	DrawingNumber          string               // 图纸号
	InstallationQuantity   int                  // 安装数量 *必填
	ImportancyLevel        base.ImportancyLevel // 重要度 *必填
	PartSpecialTreatment   string               // 零件特殊处理
	GICategory             base.GICategory      // GI类别 *必填
	SurplusStock           int                  // 剩余库存
}

func (m *MNCTemplate) UniqueOnSave() bool {
	/*
		以上四个信息不允许有重复
		并且四个当中至少有一个是有数据的
		保存模板时校验
	*/
	// ManufacturerModel      string               // 制造商的规格型号
	// ManufacturerPartNumber string               // 制造商的订货号
	// SupplierModel          string               // 供应商的规格型号
	// SupplierPartNumber     string               // 供应商的订货号
	// ^
	return true
}
