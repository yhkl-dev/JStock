package getter

import (
	"fmt"
	"jstock/src/data/mappers"
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
	userMapper *mappers.UserMapper
}

func NewUserGetterImp() *UserGetterImp {
	return &UserGetterImp{userMapper: &mappers.UserMapper{}}
}

func (s *UserGetterImp) GetUserList() (users []*userModel.User) {
	sqlMapper := s.userMapper.GetUserList()
	dbs.Orm.Exec(sqlMapper.SQL, sqlMapper.Args...).Find(&users)
	return
}

func (s *UserGetterImp) GetUserByID(id int) *result.ErrorResult {
	user := userModel.New()
	s.userMapper.GetUserList()
	db := dbs.Orm.Where("id = ?", id).Find(user)
	if db.Error != nil || db.RowsAffected == 0 {
		return result.Result(nil, fmt.Errorf("user not found, id = %d", id))
	}
	return result.Result(user, nil)
}
