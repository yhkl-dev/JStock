package services

import (
	"JStock/src/application/assembler"

	"gorm.io/gorm")


type WorkFlowMainService struct {
	AssWorkFlowAddReq *assembler.WorkFlowAddRequest
	AssWorkFlowAddReq *assembler.FlowTemplateAddRequest
	DB *gorm.DB `inject:"-"`
}

func (s *WorkFlowMainService) CreateWorkFlowTemplate(req *dto.FlowTemplateAddRequest) *dto.FlowTemplateAddResponse {
	flowTemplateModel := s.AssFlowTemplateAddReq.D2M_FlowTemplateMain(req)
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