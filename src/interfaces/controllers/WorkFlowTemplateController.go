package controllers

import (
	"JStock/src/application/dto"
	"JStock/src/application/services"
	"JStock/src/interfaces/utils"

	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type WorkFlowTemplateMainController struct {
	FlowTemplateMainSvc *services.WorkFlowTemplateService `inject:"-"`
}

func NewWorkFlowTemplateMainController() *WorkFlowTemplateMainController {
	return &WorkFlowTemplateMainController{}
}

func (s *WorkFlowTemplateMainController) CreateFlow(ctx *gin.Context) goft.Json {
	return s.FlowTemplateMainSvc.CreateWorkFlowTemplate(utils.Exec(ctx.ShouldBindJSON, &dto.FlowTemplateAddRequest{}).Unwrap().(*dto.FlowTemplateAddRequest))
}

func (s *WorkFlowTemplateMainController) ListWorkFlowTemplate(ctx *gin.Context) goft.Json {
	return s.FlowTemplateMainSvc.ListWorkFlowTemplate(utils.Exec(ctx.Bind, &dto.FlowTemplateListRequest{}).Unwrap().(*dto.FlowTemplateListRequest))
}

func (*WorkFlowTemplateMainController) Name() string {
	return "WorkFlowTemplateMainController"
}

func (s *WorkFlowTemplateMainController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/workflow/flowtemplate", s.CreateFlow)
}
