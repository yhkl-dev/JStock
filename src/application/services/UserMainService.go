package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontuser "JStock/src/domain/aggs/frontUser"
	userRoleMap "JStock/src/domain/models/userRoleMapModel"
	"JStock/src/infrastructure/GormDao"
	"JStock/src/interfaces/utils"

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
	AssUserUpdateRsp *assembler.UserUpdateResponse
	AssUserLoginReq  *assembler.UserLoginRequest
	AssUserLoginRsp  *assembler.UserLoginResponse
	DB               *gorm.DB `inject:"-"`
}

func (s *UserMainService) Login(req *dto.UserLoginRequest) *dto.UserLoginResponse {
	user := s.AssUserLoginReq.D2M_User(req)
	role := userRoleMap.New()
	repo := GormDao.NewUserMainRepo(s.DB)
	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(user, role, repo, repo2)
	err := frontUser.LoginFunc(req.Password)
	if err != nil {
		return s.AssUserLoginRsp.D2M_UserInfo(frontUser, "", err)
	}

	token, err := utils.GenerateToken(frontUser.UserMain.UserID)
	return s.AssUserLoginRsp.D2M_UserInfo(frontUser, token, err)
}

func (s *UserMainService) GetUserInfo(req *dto.UserMainRequest) *dto.UserMainResponse {
	userMainModel := s.AssUserMainReq.D2M_UserMain(req)
	userRoleMapModel := userRoleMap.New(
		userRoleMap.WithUserID(userMainModel.ID),
	)
	repo := GormDao.NewUserMainRepo(s.DB)
	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userMainModel, userRoleMapModel, repo, repo2)
	err := frontUser.QueryDetail()
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserMainRsp.D2M_UserMainInfo(frontUser)
}

func (s *UserMainService) GetUsetList(req *dto.UserListRequest) *dto.UserListResponse {
	userMainModel := s.AssUserListReq.D2M_UserList(req)
	userRoleMapModel := userRoleMap.New()
	repo := GormDao.NewUserMainRepo(s.DB)
	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
	results, err := frontuser.NewFrontUserAgg(userMainModel, userRoleMapModel, repo, repo2).
		QueryUserList(userMainModel, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserListRsp.D2M_UserListInfo(results)
}

func (s *UserMainService) CreateUser(req *dto.UserAddRequest) *dto.UserAddResponse {
	userAddModel := s.AssUserAddReq.D2M_User(req)
	userRoleMapModel := userRoleMap.New()
	repo := GormDao.NewUserMainRepo(s.DB)
	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userAddModel, userRoleMapModel, repo, repo2)
	_, err := frontUser.CreateUser(userAddModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserAddRsp.D2M_AddUserInfo(frontUser)
}

func (s *UserMainService) UpdateUser(req *dto.UserUpdateRequest) *dto.UserUpdateResponse {
	userUpdateModel := s.AssUserUpdateReq.D2M_User(req)
	userRoleMapModel := userRoleMap.New(
		userRoleMap.WithUserID(userUpdateModel.ID),
	)
	repo := GormDao.NewUserMainRepo(s.DB)
	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
	frontUser := frontuser.NewFrontUserAgg(userUpdateModel, userRoleMapModel, repo, repo2)
	err := frontUser.UpdateUser()
	if err != nil {
		panic(err.Error())
	}
	return s.AssUserUpdateRsp.D2M_UpdatedUserInfo(frontUser)
}
