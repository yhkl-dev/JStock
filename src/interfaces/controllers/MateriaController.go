package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type MaterialController struct {
	MaterialService      *services.MaterialService      `inject:"-"`
	PlantService         *services.PlantService         `inject:"-"`
	MaterialGroupService *services.MaterialGroupService `inject:"-"`
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

func (s *MaterialController) CreatePlant(ctx *gin.Context) goft.Json {
	return s.PlantService.CreatePlant(utils.Exec(ctx.ShouldBindJSON, &dto.PlantAddRequest{}).Unwrap().(*dto.PlantAddRequest))
}

func (s *MaterialController) ListPlant(ctx *gin.Context) goft.Json {
	return s.PlantService.ListPlant(utils.Exec(ctx.Bind, &dto.PlantListRequest{}).Unwrap().(*dto.PlantListRequest))
}

func (s *MaterialController) ListMaterialGroup(ctx *gin.Context) goft.Json {
	return s.MaterialGroupService.ListMaterialGroup(utils.Exec(ctx.Bind, &dto.MaterialGroupListRequest{}).Unwrap().(*dto.MaterialGroupListRequest))
}

func (s *MaterialController) ListPlantTechCode(ctx *gin.Context) goft.Json {
	return s.PlantService.ListPlantTechCode(utils.Exec(ctx.Bind, &dto.PlantTechCodeListRequest{}).Unwrap().(*dto.PlantTechCodeListRequest))
}

func (s *MaterialController) ListImportancyLevel(ctx *gin.Context) goft.Json {
	return s.MaterialService.ListImportancyLevel()
}

func (*MaterialController) Name() string {
	return "MaterialController"
}

func (s *MaterialController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/material", s.CreateMaterial).
		Handle("GET", "/material", s.ListMaterial).
		Handle("POST", "/plant", s.CreatePlant).
		Handle("GET", "/plant", s.ListPlant).
		Handle("GET", "/material/importantLevel", s.ListImportancyLevel).
		Handle("GET", "/plant/plantTech", s.ListPlantTechCode).
		Handle("GET", "/materialGroup", s.ListMaterialGroup)
}
