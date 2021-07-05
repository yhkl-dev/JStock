package services

import (
	"JStock/src/application/assembler"
	"JStock/src/application/dto"
	frontworkflowtempate "JStock/src/domain/aggs/frontWorkflowTempate"
	"JStock/src/infrastructure/GormDao"

	"gorm.io/gorm"
)

type WorkFlowItemTemplateMainService struct {
	AssItemAddReq *assembler.ItemMainRequest
	AssItemAddRsp *assembler.ItemMainResponse
	DB            *gorm.DB `inject:"-"`
}

func (s *WorkFlowItemTemplateMainService) CreateWorkFlowItem(req *dto.ItemAddRequest) *dto.ItemAddResponse {
	itemModel := s.AssItemAddReq.D2M_ItemMain(req)
	repo := GormDao.NewItemMainRepo(s.DB)
	frontItem := frontworkflowtempate.NewFrontWorkFlowItemTemplateAgg(itemModel, repo)
	err := frontItem.CreateItem(itemModel)
	if err != nil {
		panic(err.Error())
	}
	return s.AssItemAddRsp.D2M_ItemMainInfo(frontItem)
}
