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
	return s.UserMainSvc.GetUsetList(utils.Exec(ctx.Bind, &dto.UserListRequest{}).Unwrap().(*dto.UserListRequest))
}

func (*UserMainController) Name() string {
	return "UserMainController"
}

func (s *UserMainController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/user/:id", s.UserInfo).Handle("GET", "/user", s.UserList)
}
