package handlers

import (
	"fmt"
	"jstock/src/data/getter"
	"jstock/src/data/setter"
	"jstock/src/models/userModel"
	"jstock/src/result"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	Return(c)("success", "10001", getter.UserGetter.GetUserList())(OK)
}

func UserLogin(c *gin.Context) {
	Return(c)("success", "10001", gin.H{"token": "admin-token"})(OK)
}

func UserInfo(c *gin.Context) {
	Return(c)("success", "10001", gin.H{"token": "admin-token"})(OK)
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
	fmt.Println(u)
	Return(c)("success", "10001", setter.UserSetter.SaveUser(u).Unwrap())(OK)
}

func UserUpdate(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	u := userModel.New()
	result.Result(c.ShouldBindJSON(u)).Unwrap()
	u.ID = id.ID
	Return(c)("success", "10001", setter.UserSetter.UpdateUser(u).Unwrap())(OK)
}

func UserDelete(c *gin.Context) {
	id := &struct {
		ID int `uri:"id" binding:"required"`
	}{}
	result.Result(c.ShouldBindUri(id)).Unwrap()
	Return(c)("success", "10001", setter.UserSetter.DeleteUser(id.ID).Unwrap())(OK)
}
