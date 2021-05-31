package mappers

import (
	"jstock/src/models/userModel"

	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

func (s *UserMapper) AddUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Insert(user.TableName()).Columns("user_name", "userpassword").Values(user.UserName, user.UserPassword).ToSql())
}

func (s *UserMapper) UpdateUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Update(user.TableName()).Set("user_name", user.UserName).Set("userpassword", user.UserPassword).Where("id = ?", user.ID).ToSql())
}

func (s *UserMapper) GetUserList() *SQLMapper {
	return Mapper(squirrel.Select("id", "user_name").From("users").OrderBy("id desc").Limit(10).ToSql())
}

func (s *UserMapper) GetUserDetailByID(id int) *SQLMapper {
	return Mapper(squirrel.Select("id", "user_name").From("users").Where("id = ?", id).ToSql())
}
