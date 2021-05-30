package mappers

import (
	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

func (s *UserMapper) GetUserList() *SQLMapper {
	return Mapper(squirrel.Select("id", "username").From("users").OrderBy("id desc").Limit(10).ToSql())
}
