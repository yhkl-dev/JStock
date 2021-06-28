package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type UserMainService struct {
	AssUserMainReq *assembler.UserMainRequest
	AssUserMainRsp *assembler.UserMainResponse
	AssUserListReq *assembler.UserListRequest
	AssUserListRsp *assembler.UserMainResponse
	DB             *gorm.DB `inject:"-"`
}

func (s *UserMainService) GetUserInfo(req *dto.UserMainRequest) *dto.UserMainResponse {
	userMainModel := s.AssUserMainReq.D2M_UserMain(req)
	repo := GormDao.NewUserMainRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userMainModel, repo)
	err := frontUser.QueryDetail()
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserMainRsp.D2M_UserMainInfo(frontUser)
}

func (s *UserMainService) GetUsetList(req *dto.UserListRequest) []*dto.UserMainResponse {
	userMainModels := s.AssUserListReq.D2M_UserList(req)
}
