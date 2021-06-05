package handlers

import (
	"jstock/src/data/getter"
	"jstock/src/data/setter"
	"jstock/src/models/roleModel"
	"jstock/src/result"

	"github.com/gin-gonic/gin"
)

func RoleList(c *gin.Context) {
	Return(c)("success", "10001", getter.RoleGetter.GetRoleList())(OK)
	// if user.ID > 10 {
	// 	Return(c)("userlist", "10001", result.Result(test.GetInfo(user.ID)))(OK)
	// } else {
	// 	Return(c)("userlist", "10001", result.Result(test.GetInfo(user.ID)))(Error)
	// }
}

// func he,(c *gin.Context) {
// 	id := &struct {
// 		ID int `uri:"id" binding:"required"`
// 	}{}
// 	result.Result(c.ShouldBindUri(id)).Unwrap()
// 	Return(c)("success", "10001", getter.UserGetter.GetUserByID(id.ID).Unwrap())(OK)
// }

func RoleSave(c *gin.Context) {
	u := roleModel.New()
	result.Result(c.ShouldBindJSON(u)).Unwrap()
	Return(c)("success", "10001", setter.RoleSetter.SaveRole(u).Unwrap())(OK)
}

func RoleUpdate(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	r := roleModel.New()
	result.Result(c.ShouldBindJSON(r)).Unwrap()
	Return(c)("sucess", "10001", setter.RoleSetter.UpdateRole(r).Unwrap())(OK)
}

func RoleDelete(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	Return(c)("success", "10001", setter.RoleSetter.DeleteRole(id.ID).Unwrap())(OK)
}
