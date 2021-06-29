package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type RoleMainController struct {
	RoleMainSvc *services.RoleMainService `inject:"-"`
}

func NewRoleMainController() *RoleMainController {
	return &RoleMainController{}
}

func (s *RoleMainController) RoleInfo(ctx *gin.Context) goft.Json {
	return s.RoleMainSvc.GetRoleInfo(utils.Exec(ctx.ShouldBindUri, &dto.RoleMainRequest{}).Unwrap().(*dto.RoleMainRequest))
}

func (s *RoleMainController) RoleList(ctx *gin.Context) goft.Json {
	return s.RoleMainSvc.GetRoleList(utils.Exec(ctx.Bind, &dto.RoleListRequest{}).
		Unwrap().(*dto.RoleListRequest))
}

func (s *RoleMainController) CreateRole(ctx *gin.Context) goft.Json {
	return s.RoleMainSvc.CreateRole(utils.Exec(ctx.ShouldBindJSON, &dto.RoleAddRequest{}).Unwrap().(*dto.RoleAddRequest))
}

// func (s *RoleMainController) UpdateRole(ctx *gin.Context) goft.Json {
// 	return s.RoleMainSvc.UpdateRole(utils.Exec(ctx.ShouldBindJSON, &dto.UserUpdateRequest{}).Unwrap().(*dto.UserUpdateRequest))
// }

func (*RoleMainController) Name() string {
	return "RoleMainController"
}

func (s *RoleMainController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/role/:id", s.RoleInfo).
		Handle("GET", "/role", s.RoleList).
		Handle("POST", "/role", s.CreateRole)
	// Handle("PUT", "/role", s.UpdateRole)
}
