package mappers

import (
	"jstock/src/models/roleModel"
	"time"

	"github.com/Masterminds/squirrel"
)

type RoleMapper struct {
}

func (s *RoleMapper) AddRole(role *roleModel.Role) *SQLMapper {
	return Mapper(squirrel.Insert(role.TableName()).Columns("role_name", "description", "create_at").Values(role.RoleName, role.Description, time.Now()).ToSql())
}

func (s *RoleMapper) UpdateRole(role *roleModel.Role) *SQLMapper {
	return Mapper(squirrel.Update(role.TableName()).Set("role_name", role.RoleName).Set("description", role.Description).Where("id = ?", role.ID).ToSql())
}

func (s *RoleMapper) GetRoleList() *SQLMapper {
	return Mapper(squirrel.Select("id", "role_name", "description", "create_at").From("t_roles").OrderBy("id desc").Limit(10).ToSql())
}

func (s *RoleMapper) GetRoleDetailByID(id int) *SQLMapper {
	return Mapper(squirrel.Select("id", "role_name", "description", "create_at").From("t_roles").Where("id = ?", id).ToSql())
}

func (s *RoleMapper) DeleteRoleByID(id int) *SQLMapper {
	return Mapper(squirrel.Delete("t_roles").Where("id = ?", id).ToSql())
}
