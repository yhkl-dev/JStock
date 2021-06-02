package mappers

import (
	"crypto/md5"
	"fmt"
	"jstock/src/models/userModel"

	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

func (s *UserMapper) AddUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Insert(user.TableName()).Columns("user_name", "userpassword").Values(user.UserName, fmt.Sprintf("%x", md5.Sum([]byte(user.UserPassword)))).ToSql())
}

func (s *UserMapper) UpdateUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Update(user.TableName()).Set("user_name", user.UserName).Set("userpassword", fmt.Sprintf("%x", md5.Sum([]byte(user.UserPassword)))).Where("id = ?", user.ID).ToSql())
}

func (s *UserMapper) GetUserList() *SQLMapper {
	return Mapper(squirrel.Select("t_users.id ,t_users.user_name,	t_roles.role_name").From("t_users,	t_roles,	t_user_role_mapping turm ").Where("t_users.id = turm.user_id and 	t_roles.id = turm.role_id").ToSql())
		// return Mapper(squirrel.Select("id", "user_name").From("t_users").OrderBy("id desc").Limit(10).ToSql())
}

func (s *UserMapper) GetUserDetailByID(id int) *SQLMapper {
	return Mapper(squirrel.Select("id", "user_name").From("t_users").Where("id = ?", id).ToSql())
}

func (s *UserMapper) DeleteUserByID(id int) *SQLMapper {
	return Mapper(squirrel.Delete("t_users").Where("id = ?", id).ToSql())
}
