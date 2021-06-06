package mappers

import "github.com/Masterminds/squirrel"

type PurchaseTypeMapper struct {
}

func (s *PurchaseTypeMapper) GetPurchaseTypeList() *SQLMapper {
	return Mapper(squirrel.Select("id", "purchase_type", "comments").From("t_purchase_type").OrderBy("id desc").ToSql())
}
