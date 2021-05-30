package handlers

import (
	"jstock/src/data/getter"
	"jstock/src/models/userModel"
	"jstock/src/result"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	user := userModel.New()
	result.Result(c.ShouldBind(user)).Unwrap()
	getter.UserGetter.GetUserList()
	Return(c)("userlist", "10001", getter.UserGetter.GetUserList())(OK)
	// if user.ID > 10 {
	// 	Return(c)("userlist", "10001", result.Result(test.GetInfo(user.ID)))(OK)
	// } else {
	// 	Return(c)("userlist", "10001", result.Result(test.GetInfo(user.ID)))(Error)
	// }
}

func UserDetail(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	Return(c)("success", "10001", getter.UserGetter.GetUserByID(id.ID).Unwrap())(OK)
}

func UserSave(c *gin.Context) {
	u := userModel.New()
	result.Result(c.ShouldBindJSON(u)).Unwrap()
	Return(c)("success", "10001", u)(OK)
}
