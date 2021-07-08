package main

import (
	"JStock/src/infrastructure/midwares/corsmid"
	"JStock/src/infrastructure/midwares/jwtauth"
	"JStock/src/interfaces/configs"
	"JStock/src/interfaces/controllers"

	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite(corsmid.Cors()).
		Attach(jwtauth.NewTokenCheck()).
		Config(configs.NewDBConfig(),
			configs.NewUserMainServiceConfig(), 
			configs.NewRoleMainServiceConfig(),
			configs.NewUserRoleMapServiceConfig(),
			configs.NewFlowItemTemplateServiceConfig(),
			configs.NewMaterialMainServiceConfig(),
			configs.NewFlowTemplateServiceConfig()).
		Mount("v1", 
			controllers.NewUserMainControllerr(),	
			controllers.NewItemTemplateMainController(),
			controllers.NewWorkFlowTemplateMainController(),
			controllers.NewRoleMainController(), 
			controllers.NewMaterialController()).
		Launch()
}
