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

// func UserUpdate(c *gin.Context) {
// 	id := &struct {
// 		ID int `uri:"id" binding:"required"`
// 	}{}
// 	result.Result(c.ShouldBindUri(id)).Unwrap()
// 	u := userModel.New()
// 	result.Result(c.ShouldBindJSON(u)).Unwrap()
// 	u.ID = id.ID
// 	Return(c)("success", "10001", setter.UserSetter.UpdateUser(u).Unwrap())(OK)
// }

// func UserDelete(c *gin.Context) {
// 	id := &struct {
// 		ID int `uri:"id" binding:"required"`
// 	}{}
// 	result.Result(c.ShouldBindUri(id)).Unwrap()
// 	Return(c)("success", "10001", setter.UserSetter.DeleteUser(id.ID).Unwrap())(OK)
// }
