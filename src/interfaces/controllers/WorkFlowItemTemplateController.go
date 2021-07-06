package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type ItemTemplateMainController struct {
	ItemTemplateMainSvc *services.WorkFlowItemTemplateMainService `inject:"-"`
}

func NewItemTemplateMainController() *ItemTemplateMainController {
	return &ItemTemplateMainController{}
}

func (s *ItemTemplateMainController) CreateItem(ctx *gin.Context) goft.Json {
	return s.ItemTemplateMainSvc.CreateWorkFlowItem(utils.Exec(ctx.ShouldBindJSON, &dto.ItemAddRequest{}).Unwrap().(*dto.ItemAddRequest))
}

func (s *ItemTemplateMainController) UpdateItem(ctx *gin.Context) goft.Json {
	return s.ItemTemplateMainSvc.UpdateWorkFlowItem(utils.Exec(ctx.ShouldBindJSON, &dto.ItemUpdateRequest{}).Unwrap().(*dto.ItemUpdateRequest))
}


func (*ItemTemplateMainController) Name() string {
	return "ItemTemplateMainController"
}
func (s *ItemTemplateMainController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/workflow/flowtemplate/item", s.CreateItem).
		Handle("PUT", "/workflow/flowtemplate/item", s.UpdateItem)
}
