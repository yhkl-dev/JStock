package assembler

import (
	"JStock/src/application/dto"

	materialmodel "JStock/src/domain/models/materialModel"
)

type MaterialMainRequest struct{}

func (s *MaterialMainRequest) D2M_MaterialMain(dto *dto.MaterialAddRequest) *materialmodel.MaterialModel {
	m := materialmodel.New()
	m.MaterialInfo.ImportancyLevelID = dto.ImportancyLevelID
	m.MaterialInfo.MaterialGroupID = dto.MaterialGroupID
	m.MaterialInfo.RespInfo = dto.RespInfo
	m.MaterialInfo.PlantID = dto.PlantID
	m.MaterialInfo.PlantTechCodeID = dto.PlantTechCodeID
	m.MaterialInfo.HandoverType = dto.HandoverType
	m.MaterialInfo.MaterialDescriptionEN = dto.MaterialDescriptionEN
	m.MaterialInfo.MaterialDescriptionZH = dto.MaterialDescriptionZH
	m.MaterialInfo.ManufacturerName = dto.ManufacturerName
	m.MaterialInfo.ManufacturerPartNumber = dto.ManufacturerPartNumber
	m.MaterialInfo.ManufacturerModel = dto.ManufacturerModel
	m.MaterialInfo.Unit = dto.Unit
	m.MaterialInfo.Calibration = dto.Calibration
	m.MaterialInfo.Repairable = dto.Repairable
	m.MaterialInfo.Material = dto.Material
	m.MaterialInfo.CCCorCCCRelated = dto.CCCorCCCRelated
	m.MaterialInfo.PositionNumber = dto.PositionNumber
	m.MaterialInfo.MaterialMainClassification = dto.MaterialMainClassification
	m.MaterialInfo.MaterialSubClassification = dto.MaterialSubClassification
	m.MaterialInfo.ManufactureModelOld1 = dto.ManufactureModelOld1
	m.MaterialInfo.ManufactureModelOld2 = dto.ManufactureModelOld2
	m.MaterialInfo.ManufacturePNOld1 = dto.ManufacturePNOld1
	m.MaterialInfo.ManufacturePNOld2 = dto.ManufacturePNOld2
	m.MaterialInfo.Dimension = dto.Dimension
	m.MaterialInfo.MaterialSpecialTreatment = dto.MaterialSpecialTreatment
	m.MaterialInfo.MPRemark = dto.MPRemark
	m.MaterialInfo.TechRemark = dto.TechRemark
	m.MaterialInfo.SupplierName = dto.SupplierName
	m.MaterialInfo.SurplusPoint = dto.SurplusPoint
	m.MaterialInfo.InstallQty = dto.InstallQty
	return m
}

type MaterialListRequest struct{}

func (s *MaterialListRequest) D2M_MaterialMain(dto *dto.MatertialListRequest) *materialmodel.MaterialModel {
	m := materialmodel.New()
	m.MaterialNumber = dto.MaterialNumber
	m.MaterialInfo.ImportancyLevelID = dto.ImportancyLevelID
	m.MaterialInfo.MaterialGroupID = dto.MaterialGroupID
	m.MaterialInfo.PlantID = dto.PlantID
	// m.MaterialInfo.PlantTechCodeID = []int{dto.PlantTechCodeID}
	m.MaterialInfo.HandoverType = dto.HandoverType
	m.MaterialInfo.MaterialDescriptionEN = dto.MaterialDescriptionEN
	m.MaterialInfo.MaterialDescriptionZH = dto.MaterialDescriptionZH
	return m
}
