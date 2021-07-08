package assembler

import (
	"JStock/src/application/dto"

	frontmaterial "JStock/src/domain/aggs/frontMaterial"
	materialmodel "JStock/src/domain/models/materialModel"
)

type MaterialMainResponse struct{}

func (s *MaterialMainResponse) D2M_MaterialMain(agg *frontmaterial.FrontMaterialAgg) *dto.MaterialMainResponse {
	m := &dto.MaterialMainResponse{}

	m.ImportancyLevelID = agg.MaterialMain.MaterialInfo.ImportancyLevelID
	m.ImportancyLevelName = agg.MaterialMain.MaterialInfo.ImportancyLevelName
	m.MaterialDescriptionEN = agg.MaterialMain.MaterialInfo.MaterialDescriptionEN
	m.MaterialDescriptionZH = agg.MaterialMain.MaterialInfo.MaterialDescriptionZH
	m.PlatnName = agg.MaterialMain.MaterialInfo.PlatnName
	m.PlantID = agg.MaterialMain.MaterialInfo.PlantID
	m.ManufacturerName = agg.MaterialMain.MaterialInfo.ManufacturerName
	m.ManufacturerPartNumber = agg.MaterialMain.MaterialInfo.ManufacturerPartNumber
	m.ManufacturerModel = agg.MaterialMain.MaterialInfo.ManufacturerModel
	m.Unit = agg.MaterialMain.MaterialInfo.Unit
	m.Calibration = agg.MaterialMain.MaterialInfo.Calibration
	m.Repairable = agg.MaterialMain.MaterialInfo.Repairable
	m.Material = agg.MaterialMain.MaterialInfo.Material
	m.CCCorCCCRelated = agg.MaterialMain.MaterialInfo.CCCorCCCRelated
	m.PositionNumber = agg.MaterialMain.MaterialInfo.PositionNumber
	m.MaterialMainClassification = agg.MaterialMain.MaterialInfo.MaterialMainClassification
	m.MaterialSubClassification = agg.MaterialMain.MaterialInfo.MaterialSubClassification
	m.ManufactureModelOld1 = agg.MaterialMain.MaterialInfo.ManufactureModelOld1
	m.ManufactureModelOld2 = agg.MaterialMain.MaterialInfo.ManufactureModelOld2
	m.ManufacturePNOld1 = agg.MaterialMain.MaterialInfo.ManufacturePNOld1
	m.ManufacturePNOld2 = agg.MaterialMain.MaterialInfo.ManufacturePNOld2
	m.Dimension = agg.MaterialMain.MaterialInfo.Dimension
	m.MaterialSpecialTreatment = agg.MaterialMain.MaterialInfo.MaterialSpecialTreatment
	m.MPRemark = agg.MaterialMain.MaterialInfo.MPRemark
	m.TechRemark = agg.MaterialMain.MaterialInfo.TechRemark
	m.SupplierName = agg.MaterialMain.MaterialInfo.SupplierName
	m.SurplusPoint = agg.MaterialMain.MaterialInfo.SurplusPoint
	m.InstallQty = agg.MaterialMain.MaterialInfo.InstallQty

	return m
}

type MaterialListResponse struct{}

func (s *MaterialListResponse) D2M_MaterialList(agg interface{}) *dto.MaterialListResponse {
	materialList := make(dto.MaterialListResponse, 0)
	resMaterialList := agg.([]*materialmodel.MaterialModel)
	for _, m := range resMaterialList {
		mInfo := &dto.MaterialMainResponse{}
		mInfo.Id = m.ID
		mInfo.MaterialNumber = m.MaterialNumber
		mInfo.ImportancyLevelID = m.MaterialInfo.ImportancyLevelID
		mInfo.ImportancyLevelName = m.MaterialInfo.ImportancyLevelName
		mInfo.MaterialGroupID = m.MaterialInfo.MaterialGroupID
		mInfo.MaterialGroupName = m.MaterialInfo.MaterialGroupName
		mInfo.RespInfo = m.MaterialInfo.RespInfo
		mInfo.PlantID = m.MaterialInfo.PlantID
		mInfo.PlatnName = m.MaterialInfo.PlatnName
		mInfo.HandoverType = m.MaterialInfo.HandoverType
		mInfo.MaterialDescriptionEN = m.MaterialInfo.MaterialDescriptionEN
		mInfo.MaterialDescriptionZH = m.MaterialInfo.MaterialDescriptionZH
		mInfo.ManufacturerName = m.MaterialInfo.ManufacturerName
		mInfo.ManufacturerPartNumber = m.MaterialInfo.ManufacturerPartNumber
		mInfo.ManufacturerModel = m.MaterialInfo.ManufacturerModel
		mInfo.Unit = m.MaterialInfo.Unit
		mInfo.Calibration = m.MaterialInfo.Calibration
		mInfo.Repairable = m.MaterialInfo.Repairable
		mInfo.Material = m.MaterialInfo.Material
		mInfo.CCCorCCCRelated = m.MaterialInfo.CCCorCCCRelated
		mInfo.PositionNumber = m.MaterialInfo.PositionNumber
		mInfo.MaterialMainClassification = m.MaterialInfo.MaterialMainClassification
		mInfo.MaterialSubClassification = m.MaterialInfo.MaterialSubClassification
		mInfo.ManufactureModelOld1 = m.MaterialInfo.ManufactureModelOld1
		mInfo.ManufactureModelOld2 = m.MaterialInfo.ManufactureModelOld2
		mInfo.ManufacturePNOld1 = m.MaterialInfo.ManufacturePNOld1
		mInfo.ManufacturePNOld2 = m.MaterialInfo.ManufacturePNOld2
		mInfo.Dimension = m.MaterialInfo.Dimension
		mInfo.MaterialSpecialTreatment = m.MaterialInfo.MaterialSpecialTreatment
		mInfo.MPRemark = m.MaterialInfo.MPRemark
		mInfo.TechRemark = m.MaterialInfo.TechRemark
		mInfo.SupplierName = m.MaterialInfo.SupplierName
		mInfo.SurplusPoint = m.MaterialInfo.SurplusPoint
		mInfo.InstallQty = m.MaterialInfo.InstallQty

		materialList = append(materialList, *mInfo)
	}
	return &materialList
}
