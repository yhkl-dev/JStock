package configs

import (
	"JStock/src/application/services"
)

type UserRoleMapServiceConfig struct{}

func NewUserRoleMapServiceConfig() *UserRoleMapServiceConfig {
	return &UserRoleMapServiceConfig{}
}

func (s *UserRoleMapServiceConfig) UserRoleMapService() *services.UserRoleMapService {
	return &services.UserRoleMapService{}
}
