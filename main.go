package main

import (
	"jstock/src/common"
	"jstock/src/dbs"
	"jstock/src/handlers"
	_ "jstock/src/validators"

	"github.com/gin-gonic/gin"
)

func main() {
	dbs.InitDB()
	r := gin.New()

	r.Use(common.ErrorHandler())
	r.GET("/users", handlers.UserList)
	r.GET("/users/:id", handlers.UserDetail)
	r.POST("/users", handlers.UserSave)
	r.PUT("/users/:id", handlers.UserUpdate)
	r.DELETE("/users/:id", handlers.UserDelete)
	r.GET("/roles", handlers.RoleList)
	r.POST("/roles", handlers.RoleSave)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
