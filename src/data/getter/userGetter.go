package getter

import (
	"fmt"
	"jstock/src/dbs"
	"jstock/src/models/userModel"
	"jstock/src/result"
)

var UserGetter IUserGetter

type IUserGetter interface {
	GetUserList() []*userModel.User
	GetUserByID(id int) *result.ErrorResult
}

func init() {
	UserGetter = NewUserGetterImp()
}

type UserGetterImp struct {
}

func NewUserGetterImp() *UserGetterImp {
	return &UserGetterImp{}
}

func (s *UserGetterImp) GetUserList() (users []*userModel.User) {
	dbs.Orm.Find(&users)
	return
}

func (s *UserGetterImp) GetUserByID(id int) *result.ErrorResult {
	user := userModel.New()
	db := dbs.Orm.Where("id = ?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("user not found, id = %d", id))
	}
	return result.Result(user, nil)
}
