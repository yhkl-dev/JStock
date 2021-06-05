package mappers

import (
	"crypto/md5"
	"fmt"
	"jstock/src/models/userModel"
	"time"

	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

func (s *UserMapper) AddUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Insert(user.TableName()).Columns("user_name_zh", "user_name_en", "user_mobile_phone", "user_telephone", "user_mail", "remark", "create_at", "userpassword").Values(user.UserNameZh, user.UserNameEn, user.UserMobilePhone, user.UserTelePhone, user.UserMail, user.Remark, time.Now(), fmt.Sprintf("%x", md5.Sum([]byte(user.UserPassword)))).ToSql())
}

func (s *UserMapper) UpdateUser(user *userModel.User) *SQLMapper {
	return Mapper(squirrel.Update(user.TableName()).
		Set("user_name_zh", user.UserNameZh).
		Set("user_name_en", user.UserNameEn).
		Set("user_mobile_phone", user.UserMobilePhone).
		Set("user_telephone", user.UserTelePhone).
		Set("user_mail", user.UserMail).
		Set("remark", user.Remark).
		Set("userpassword", fmt.Sprintf("%x", md5.Sum([]byte(user.UserPassword)))).Where("id = ?", user.ID).ToSql())
}

func (s *UserMapper) GetUserList() *SQLMapper {
	return Mapper(squirrel.Select("id, user_id ,user_name_zh, user_name_en, user_mobile_phone, user_telephone, user_mail, remark").From("t_users").ToSql())
	// return Mapper(squirrel.Select("id", "user_name").From("t_users").OrderBy("id desc").Limit(10).ToSql())
}

func (s *UserMapper) GetUserDetailByID(id int) *SQLMapper {
	return Mapper(squirrel.Select("user_id ,user_name_zh, user_name_en, user_mobile_phone, user_telephone, user_mail, remark").From("t_users").Where("id = ?", id).ToSql())
}

func (s *UserMapper) DeleteUserByID(id int) *SQLMapper {
	return Mapper(squirrel.Delete("t_users").Where("id = ?", id).ToSql())
}

func (s *UserMapper) DeteleUserRole(id int) *SQLMapper {
	return Mapper(squirrel.Delete("t_user_role_mapping").Where("user_id = ?", id).ToSql())
}
