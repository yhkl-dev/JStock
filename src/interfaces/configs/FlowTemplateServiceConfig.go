package configs

import "JStock/src/application/services"

type FlowTemplateServiceConfig struct{}

func NewFlowTemplateServiceConfig() *FlowTemplateServiceConfig {
	return &FlowTemplateServiceConfig{}
}

func (s *FlowTemplateServiceConfig) FlowTemplateConfigService() *services.WorkFlowTemplateService {
	return &services.WorkFlowTemplateService{}
}

type FlowItemTemplateServiceConfig struct{}

func NewFlowItemTemplateServiceConfig() *FlowItemTemplateServiceConfig {
	return &FlowItemTemplateServiceConfig{}
}

func (s *FlowItemTemplateServiceConfig) FlowItemTemplateService() *services.WorkFlowItemTemplateMainService {
	return &services.WorkFlowItemTemplateMainService{}
}
