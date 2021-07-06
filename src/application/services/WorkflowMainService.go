package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontworkflowtempate "JStock/src/domain/aggs/frontWorkflowTempate"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type WorkFlowMainService struct {
	AssWorkFlowAddReq *assembler.WorkFlowAddRequest
	AssWorkFlowAddRsp *assembler.WorkFlowAddRequest
	DB                *gorm.DB `inject:"-"`
}

func (s *WorkFlowMainService) CreateWorkFlowTemplate(req *dto.FlowTemplateAddRequest) *dto.FlowTemplateAddResponse {
	flowTemplateModel := s.AssWorkFlowAddReq.D2M_FlowTemplateMain(req)
	ItemModel := workflowtemplate.NewWorkFlowItemTempate()
	repo1 := GormDao.NewIFlowTemplateRepo(s.DB)
	repo2 := GormDao.NewItemMainRepo(s.DB)
	frontFlowTemplate := frontworkflowtempate.NewFrontWorkFlowTemplateAgg(flowTemplateModel, ItemModel, repo1, repo2)
	err := frontFlowTemplate.UpdateFlowTemplate(flowTemplateModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssFlowTemplateAddRsp.D2M_FlowTemplateInfo(frontFlowTemplate)
}
