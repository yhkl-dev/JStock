package assembler

import (
	"JStock/src/application/dto"
	materialmodel "JStock/src/domain/models/materialModel"
)

type MaterialGroupListRequest struct{}

func (s *MaterialGroupListRequest) D2M_MaterialGroup(dto *dto.MaterialGroupListRequest) *materialmodel.MaterialGroup {
	return materialmodel.NewMaterialGroup()
}
