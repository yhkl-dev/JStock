package main

import (
	"JStock/src/infrastructure/midwares/jwtauth"
	"JStock/src/interfaces/configs"
	"JStock/src/interfaces/controllers"

	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().Attach(jwtauth.NewTokenCheck()).
		Config(configs.NewDBConfig(),
			configs.NewUserMainServiceConfig(),
			configs.NewRoleMainServiceConfig(),
			configs.NewUserRoleMapServiceConfig()).
		Mount("v1", controllers.NewUserMainControllerr(),
			controllers.NewRoleMainController()).
		Launch()
}
