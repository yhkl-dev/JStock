package services

import (
	"JStock/src/application/assembler"

	"gorm.io/gorm"
)

type MaterialService struct {
	AssMaterialMainReq *assembler.MaterialMainRequest
	AssRoleMainReq *assembler.RoleMainRequest
	AssRoleMainRsp *assembler.RoleMainResponse
	AssRoleListReq *assembler.RoleListRequest
	AssRoleListRsp *assembler.RoleListResponse
	AssRoleAddReq  *assembler.RoleAddRequest
	AssRoleAddRsp  *assembler.RoleAddResponse
	DB             *gorm.DB `inject:"-"`
}


func (s *MaterialService) GetMaterialInfo() error {
	return nil 
}