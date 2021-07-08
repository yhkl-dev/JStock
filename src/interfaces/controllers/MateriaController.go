package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type MaterialController struct {
	MaterialService *services.MaterialService `inject:"-"`
}

func NewMaterialController() *MaterialController {
	return &MaterialController{}
}

func (s *MaterialController) CreateMaterial(ctx *gin.Context) goft.Json {
	return s.MaterialService.CreateMaterial(utils.Exec(ctx.ShouldBindJSON, &dto.MaterialAddRequest{}).Unwrap().(*dto.MaterialAddRequest))
}

func (s *MaterialController) ListMaterial(ctx *gin.Context) goft.Json {
	return s.MaterialService.ListMaterial(utils.Exec(ctx.Bind, &dto.MatertialListRequest{}).Unwrap().(*dto.MatertialListRequest))
}

func (*MaterialController) Name() string {
	return "MaterialController"
}

func (s *MaterialController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/material", s.CreateMaterial).
		Handle("GET", "/material", s.ListMaterial)
}
