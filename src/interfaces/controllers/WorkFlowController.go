package controllers

import "JStock/src/application/services"

type WorkFlowMainController struct {
	UserMainSvc *services.WorkFlowMainService `inject:"-"`
}

func NewWorkFlowMainController() *WorkFlowMainController {
	return &WorkFlowMainController{}
}