package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontworkflow "JStock/src/domain/aggs/frontWorkflow"
	workflowtemplate "JStock/src/domain/models/workFlowTemplate"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type WorkFlowMainService struct {
	AssWorkFlowAddReq *assembler.WorkFlowAddRequest
	AssWorkFlowAddRsp *assembler.WorkFlowAddResponse
	DB                *gorm.DB `inject:"-"`
}

func (s *WorkFlowMainService) CreateWorkFlow(req *dto.WorkFlowAddRequest) *dto.WorkFlowAddResponse {
	workflow := s.AssWorkFlowAddReq.D2M_WorkFlow(req)
	template := workflowtemplate.NewWorkFlowTemplate()
	repo1 := GormDao.NewWorkFlowRepoo(s.DB)
	repo2 := GormDao.NewIFlowTemplateRepo(s.DB)
	f := frontworkflow.NewFrontWorkFlowAgg(workflow, template, repo1, repo2)
	err := f.CreateWorkFlow(workflow)
	if err != nil {
		panic(err.Error())
	}
	return s.AssWorkFlowAddRsp.D2M_WorkFlow(f)
}
