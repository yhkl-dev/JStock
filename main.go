package main

import (
	"JStock/src/interfaces/configs"
	"JStock/src/interfaces/controllers"

	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().
		Config(configs.NewDBConfig(),
			configs.NewUserMainServiceConfig(),
			configs.NewRoleMainServiceConfig()).
		Mount("v1", controllers.NewUserMainControllerr(),
			controllers.NewRoleMainController()).
		Launch()
}
