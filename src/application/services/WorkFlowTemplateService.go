package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontworkflowtempate "JStock/src/domain/aggs/frontWorkflowTempate"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type WorkFlowTemplateService struct {
	AssFlowTemplateAddReq  *assembler.FlowTemplateAddRequest
	AssFlowTemplateAddRsp  *assembler.FlowTemplateAddResponse
	AssFlowTemplateListReq *assembler.FlowTemplateListRequest
	AssFlowTemplateListRsp *assembler.FlowTemplateListResponse
	DB                     *gorm.DB `inject:"-"`
}

func (s *WorkFlowTemplateService) CreateWorkFlowTemplate(req *dto.FlowTemplateAddRequest) *dto.FlowTemplateAddResponse {
	flowTemplateModel := s.AssFlowTemplateAddReq.D2M_FlowTemplateMain(req)
	ItemModel := workflowtemplate.NewWorkFlowItemTempate()
	repo1 := GormDao.NewIFlowTemplateRepo(s.DB)
	repo2 := GormDao.NewItemMainRepo(s.DB)
	frontFlowTemplate := frontworkflowtempate.NewFrontWorkFlowTemplateAgg(flowTemplateModel, ItemModel, repo1, repo2)
	err := frontFlowTemplate.CreateFlowTemplate(flowTemplateModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssFlowTemplateAddRsp.D2M_FlowTemplateInfo(frontFlowTemplate)
}

func (s *WorkFlowTemplateService) ListWorkFlowTemplate(req *dto.FlowTemplateListRequest) *dto.FlowTemplateListResponse {
	flowModel := s.AssFlowTemplateListReq.D2M_FlowTemplateList(req)
	ItemModel := workflowtemplate.NewWorkFlowItemTempate()
	repo1 := GormDao.NewIFlowTemplateRepo(s.DB)
	repo2 := GormDao.NewItemMainRepo(s.DB)
	frontFlowTemplate := frontworkflowtempate.NewFrontWorkFlowTemplateAgg(flowModel, ItemModel, repo1, repo2)
	results, err := frontFlowTemplate.ListFlowTemplate(flowModel, req.Page, req.PageSize)
	if err != nil {
		panic(err.Error())
	}
	return s.AssFlowTemplateListRsp.D2M_FlowTemplateLIst(results)
}
