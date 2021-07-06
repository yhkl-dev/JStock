package assembler

import (
	"JStock/src/application/dto"

	frontmaterial "JStock/src/domain/aggs/frontMaterial"
)


type MaterialMainRequest struct{}


func (s *MaterialMainRequest) D2M_MaterialMain(uagg *frontmaterial.FrontMaterialAgg) *dto.MaterialMainResponse {
	m := &dto.MaterialMainResponse{}
	return m
}
