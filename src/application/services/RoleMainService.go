package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontrole "JStock/src/domain/aggs/frontRole"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type RoleMainService struct {
	AssRoleMainReq *assembler.RoleMainRequest
	AssRoleMainRsp *assembler.RoleMainResponse
	AssRoleListReq *assembler.RoleListRequest
	AssRoleListRsp *assembler.RoleListResponse
	AssRoleAddReq  *assembler.RoleAddRequest
	AssRoleAddRsp  *assembler.RoleAddResponse
	DB             *gorm.DB `inject:"-"`
}

func (s *RoleMainService) GetRoleInfo(req *dto.RoleMainRequest) *dto.RoleMainResponse {
	roleMainModel := s.AssRoleMainReq.D2M_RoleMain(req)
	repo := GormDao.NewRoleMainRepo(s.DB)
	frontRole := frontrole.NewFrontRoleAgg(roleMainModel, repo)
	err := frontRole.QueryDetail()
	if err != nil {
		panic(err.Error())
	}
	return s.AssRoleMainRsp.D2M_RoleMainInfo(frontRole)
}

func (s *RoleMainService) GetRoleList(req *dto.RoleListRequest) *dto.RoleListResponse {
	roleMainModel := s.AssRoleListReq.D2M_RoleList(req)
	repo := GormDao.NewRoleMainRepo(s.DB)
	results, err := frontrole.NewFrontRoleAgg(roleMainModel, repo).
		QueryRoleList(roleMainModel, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssRoleListRsp.D2M_RoleListInfo(results)
}

func (s *RoleMainService) CreateRole(req *dto.RoleAddRequest) *dto.RoleAddResponse {
	roleAddModel := s.AssRoleAddReq.D2M_Role(req)
	repo := GormDao.NewRoleMainRepo(s.DB)
	frontRole := frontrole.NewFrontRoleAgg(roleAddModel, repo)
	err := frontRole.CreateRole(roleAddModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssRoleAddRsp.D2M_AddRoleInfo(frontRole)
}
