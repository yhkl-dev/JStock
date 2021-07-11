package services

import (
	"JStock/src/application/assembler"

	"gorm.io/gorm"
)

type MaterialPlaningService struct {
	AssMaterialPlaningReq *assembler.MaterialGroupListRequest
	AssMaterialPlaningRsp *assembler.MaterialGroupListResponse
	DB                    *gorm.DB `inject:"-"`
}
