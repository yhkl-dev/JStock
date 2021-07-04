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

// func (s *UserMainService) CreateUser(req *dto.UserAddRequest) *dto.UserAddResponse {
// 	userAddModel := s.AssUserAddReq.D2M_User(req)
// 	userRoleMapModel := userRoleMap.New()
// 	repo := GormDao.NewUserMainRepo(s.DB)
// 	repo2 := GormDao.NewUserRoleMapRepo(s.DB)
// 	frontUser := frontuser.NewFrontUserAgg(userAddModel, userRoleMapModel, repo, repo2)
// 	_, err := frontUser.CreateUser(userAddModel)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return s.AssUserAddRsp.D2M_AddUserInfo(frontUser)
// }
