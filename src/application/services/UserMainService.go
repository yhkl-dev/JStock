package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type UserMainService struct {
	AssUserMainReq   *assembler.UserMainRequest
	AssUserMainRsp   *assembler.UserMainResponse
	AssUserListReq   *assembler.UserListRequest
	AssUserListRsp   *assembler.UserListResponse
	AssUserAddReq    *assembler.UserAddRequest
	AssUserAddRsp    *assembler.UserAddResponse
	AssUserUpdateReq *assembler.UserUpdateRequest
	AssUserUpdateRsp *assembler.UserAddResponse
	DB               *gorm.DB `inject:"-"`
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

func (s *UserMainService) GetUsetList(req *dto.UserListRequest) *dto.UserListResponse {
	userMainModel := s.AssUserListReq.D2M_UserList(req)
	repo := GormDao.NewUserMainRepo(s.DB)
	results, err := frontuser.NewFrontUserAgg(userMainModel, repo).QueryUserList(userMainModel, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserListRsp.D2M_UserListInfo(results)
}

func (s *UserMainService) CreateUser(req *dto.UserAddRequest) *dto.UserAddResponse {
	userAddModel := s.AssUserAddReq.D2M_User(req)
	repo := GormDao.NewUserMainRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userAddModel, repo)
	err := frontUser.CreateUser(userAddModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserAddRsp.D2M_AddUserInfo(frontUser)
}

func (s *UserMainService) UpdateUser(req *dto.UserUpdateRequest) *dto.UserMainResponse {
	userUpdateModel := s.AssUserUpdateReq.D2M_User(req)
	repo := GormDao.NewUserMainRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userUpdateModel, repo)
	err := frontUser.UpdateUser()
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserMainRsp.D2M_UserMainInfo(frontUser)
}
