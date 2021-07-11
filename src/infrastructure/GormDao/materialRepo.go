package GormDao

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"
	"JStock/src/interfaces/utils"
	"time"

	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
)

type MaterialRepo struct {
	db *gorm.DB
}

func NewMaterialRepo(db *gorm.DB) *MaterialRepo {
	return &MaterialRepo{db: db}
}

func (s *MaterialRepo) New(model repos.IModel) error {

	m := model.(*materialmodel.MaterialModel)
	m.MaterialNumber = utils.GenerateMaterialID()

	sql, args, _ := squirrel.Insert("t_material").Columns(
		"material_number",
		"importancy_level_id",
		"material_group_id",
		"plant_id",
		"handover_type",
		"material_description_en",
		"material_description_zh",
		"manufacturer_name",
		"manufacturer_part_number",
		"manufacturer_model",
		"unit",
		"calibration",
		"repairable",
		"material",
		"ccc_or_ccc_related",
		"position_number",
		"material_main_classification",
		"material_sub_classification",
		"manufacture_model_old1",
		"manufacture_model_old2",
		"manufacture_pn_old1",
		"manufacture_pn_old2",
		"dimension",
		"material_special_treatment",
		"mpr_remark",
		"tech_remark",
		"supplier_name",
		"surplus_point",
		"install_qty",
		"create_at").
		Values(m.MaterialNumber,
			m.MaterialInfo.ImportancyLevelID,
			m.MaterialInfo.MaterialGroupID,
			m.MaterialInfo.PlantID,
			m.MaterialInfo.HandoverType,
			m.MaterialInfo.MaterialDescriptionEN,
			m.MaterialInfo.MaterialDescriptionZH,
			m.MaterialInfo.ManufacturerName,
			m.MaterialInfo.ManufacturerPartNumber,
			m.MaterialInfo.ManufacturerModel,
			m.MaterialInfo.Unit,
			m.MaterialInfo.Calibration,
			m.MaterialInfo.Repairable,
			m.MaterialInfo.Material,
			m.MaterialInfo.CCCorCCCRelated,
			m.MaterialInfo.PositionNumber,
			m.MaterialInfo.MaterialMainClassification,
			m.MaterialInfo.MaterialSubClassification,
			m.MaterialInfo.ManufactureModelOld1,
			m.MaterialInfo.ManufactureModelOld2,
			m.MaterialInfo.ManufacturePNOld1,
			m.MaterialInfo.ManufacturePNOld2,
			m.MaterialInfo.Dimension,
			m.MaterialInfo.MaterialSpecialTreatment,
			m.MaterialInfo.MPRemark,
			m.MaterialInfo.TechRemark,
			m.MaterialInfo.SupplierName,
			m.MaterialInfo.SurplusPoint,
			m.MaterialInfo.InstallQty,
			time.Now(),
		).ToSql()

	err := s.db.Exec(sql, args...).Error
	if err != nil {
		return err
	}

	for _, user := range m.MaterialInfo.RespInfo {
		insert_user_sql, args, _ :=
			squirrel.Insert("t_material_user_mapping").
				Columns("material_id", "user_id").
				Values(m.ID, user).
				ToSql()
		if err := s.db.Exec(insert_user_sql, args...).Error; err != nil {
			return err
		}
	}
	for _, techCode := range m.MaterialInfo.PlantTechCodeID {
		insert_tech_sql, args, _ := squirrel.
			Insert("t_material_plant_tech_code_mapping").
			Columns("material_id", "plant_tech_code_id").
			Values(m.ID, techCode).
			ToSql()
		if err := s.db.Exec(insert_tech_sql, args...).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *MaterialRepo) Load(model repos.IModel) error {
	return s.db.Table("t_material").Where("id = ?", model.(*materialmodel.MaterialModel).ID).First(model).Error
}

func (s *MaterialRepo) QueryImportancyLevelList() (interface{}, error) {
	res := []*materialmodel.ImportancyLevel{}
	err := s.db.Table("t_importancy_level").Find(&res).Error
	return res, err
}

func (s *MaterialRepo) QueryMaterialList(model repos.IModel, page, pageSize int) (interface{}, error) {
	m := model.(*materialmodel.MaterialModel)
	res := []*materialmodel.MaterialModel{}
	sqlObj := squirrel.Select(`
			tm.id,
			tm.material_number,
			tm.handover_type,
			tm.material_description_en,
			tm.material_description_zh,
			tm.manufacturer_name,
			tm.manufacturer_part_number,
			tm.manufacturer_model,
			tm.unit,
			tm.calibration,
			tm.repairable,
			tm.material,
			tm.ccc_or_ccc_related,
			tm.position_number,
			tm.material_main_classification,
			tm.material_sub_classification,
			tm.manufacture_model_old1,
			tm.manufacture_model_old2,
			tm.manufacture_pn_old1,
			tm.manufacture_pn_old2,
			tm.dimension,
			tm.material_special_treatment,
			tm.mpr_remark,
			tm.tech_remark,
			tm.importancy_level_id importancy_level_id,
			til.level_name importancy_level_name,
			tm.material_group_id,
			tmg.group_name material_group_name,
			tm.supplier_name,
			tm.surplus_point,
			tm.install_qty,
			tm.plant_id,
			tp.plant_name plant_name`).
		From(`
			t_material tm,
			t_material_group tmg ,
			t_importancy_level til ,
			t_plant tp
		`).Where(`tm.material_group_id = tmg.id
			and tm.importancy_level_id = til.id
			and tm.plant_id = tp.id`)
	if m.ID != 0 {
		sqlObj = sqlObj.Where("tm.id = ?", m.ID)
	}
	if m.MaterialInfo.MaterialGroupID != 0 {
		sqlObj = sqlObj.Where("tm.material_group_id = ?", m.MaterialInfo.MaterialGroupID)
	}
	if m.MaterialNumber != "" {
		sqlObj = sqlObj.Where("tm.material_number = ?", m.MaterialNumber)
	}
	if m.MaterialInfo.MaterialDescriptionEN != "" {
		sqlObj = sqlObj.Where("tm.material_description_en = ?", m.MaterialInfo.MaterialDescriptionEN)
	}
	if m.MaterialInfo.MaterialDescriptionZH != "" {
		sqlObj = sqlObj.Where("tm.material_description_en = ?", m.MaterialInfo.MaterialDescriptionZH)
	}
	if m.MaterialInfo.HandoverType != "" {
		sqlObj = sqlObj.Where("tm.handover_type = ?", m.MaterialInfo.HandoverType)
	}
	if m.MaterialInfo.PlantID != 0 {
		sqlObj = sqlObj.Where("tm.plant_id = ?", m.MaterialInfo.PlantID)
	}
	if m.MaterialInfo.ImportancyLevelID != 0 {
		sqlObj = sqlObj.Where("tm.importancy_level_id = ?", m.MaterialInfo.ImportancyLevelID)
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 0
	}
	sql, args, _ := sqlObj.Offset(uint64(page)).Limit(uint64(pageSize)).ToSql()
	err := s.db.Raw(sql, args...).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
