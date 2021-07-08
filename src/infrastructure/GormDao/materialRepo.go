package GormDao

import (
	materialmodel "JStock/src/domain/models/materialModel"
	"JStock/src/domain/models/repos"
	"fmt"

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
	return s.db.Table("t_material").Create(model).Error
}

func (s *MaterialRepo) Load(model repos.IModel) error {
	return s.db.Table("t_material").Where("id = ?", model.(*materialmodel.MaterialModel).ID).First(model).Error
}

func (s *MaterialRepo) QueryMaterialList(model repos.IModel, page, pageSize int) (interface{}, error) {
	m := model.(*materialmodel.MaterialModel)
	fmt.Println("MaterialNumber", m.MaterialNumber)
	// var res = make([]*materialmodel.MaterialModel, 0)
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
			tmg.group_name plant_name`).
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
	sql, args, _ := sqlObj.Offset(uint64(page)).Limit(uint64(pageSize)).ToSql()
	fmt.Println(sql)
	err := s.db.Raw(sql, args...).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
