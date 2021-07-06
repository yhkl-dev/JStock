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

func (s *WorkFlowTemplateMainController) UpdateWorkFlowtemplate(ctx *gin.Context) goft.Json {
	return s.FlowTemplateMainSvc.UpdateWorkFlowTemplate(utils.Exec(ctx.ShouldBindJSON, &dto.FlowTemplateUpdateRequest{}).Unwrap().(*dto.FlowTemplateUpdateRequest))
}

func (*WorkFlowTemplateMainController) Name() string {
	return "WorkFlowTemplateMainController"
}

func (s *WorkFlowTemplateMainController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/workflow/flowtemplate", s.CreateFlow).
		Handle("GET", "/workflow/flowtemplate", s.ListWorkFlowTemplate).
		Handle("PUT", "/workflow/flowtemplate", s.UpdateWorkFlowtemplate)
}
