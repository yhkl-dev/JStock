package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserMainController struct {
	UserMainSvc *services.UserMainService `inject:"-"`
}

func NewUserMainControllerr() *UserMainController {
	return &UserMainController{}
}

func (s *UserMainController) UserInfo(ctx *gin.Context) goft.Json {
	return s.UserMainSvc.GetUserInfo(utils.Exec(ctx.ShouldBindUri, &dto.UserMainRequest{}).Unwrap().(*dto.UserMainRequest))
}

func (s *UserMainController) UserList(ctx *gin.Context) goft.Json {
	return s.UserMainSvc.GetUsetList(utils.Exec(ctx.Bind, &dto.UserListRequest{}).
		Unwrap().(*dto.UserListRequest))
}

func (s *UserMainController) CreateUser(ctx *gin.Context) goft.Json {
	return s.UserMainSvc.CreateUser(utils.Exec(ctx.ShouldBindJSON, &dto.UserAddRequest{}).Unwrap().(*dto.UserAddRequest))
}

func (s *UserMainController) UpdateUser(ctx *gin.Context) goft.Json {
	return s.UserMainSvc.UpdateUser(utils.Exec(ctx.ShouldBindJSON, &dto.UserUpdateRequest{}).Unwrap().(*dto.UserUpdateRequest))
}

func (s *UserMainController) Login(ctx *gin.Context) goft.Json {
	return s.UserMainSvc.Login(utils.Exec(ctx.Bind, &dto.UserLoginRequest{}).Unwrap().(*dto.UserLoginRequest))
}

func (*UserMainController) Name() string {
	return "UserMainController"
}

func (s *UserMainController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user/:id", s.UserInfo).
		Handle("GET", "/user", s.UserList).
		Handle("POST", "/user", s.CreateUser).
		Handle("PUT", "/user", s.UpdateUser).
		Handle("POST", "/user/login", s.Login)
}
