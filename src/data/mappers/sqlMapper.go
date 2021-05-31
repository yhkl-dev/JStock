package mappers

import (
	"jstock/src/dbs"

	"gorm.io/gorm"
)

type SQLMapper struct {
	SQL  string
	Args []interface{}
	db   *gorm.DB
}

func (s *SQLMapper) setDB(db *gorm.DB) {
	s.db = db
}

type SQLMappers []*SQLMapper

func Mappers(sqlMappers ...*SQLMapper) (list SQLMappers) {
	list = sqlMappers
	return
}

func (s SQLMappers) apply(tx *gorm.DB) {
	for _, sql := range s {
		sql.setDB(tx)
	}
}

func (s SQLMappers) Exec(f func() error) error {
	return dbs.Orm.Transaction(func(tx *gorm.DB) error {
		s.apply(tx)
		return f()
	})
}

func NewSQLMapper(sql string, args []interface{}) *SQLMapper {
	return &SQLMapper{SQL: sql, Args: args}
}

func (s *SQLMapper) Query() *gorm.DB {
	if s.db != nil {
		return s.db.Raw(s.SQL, s.Args...)
	}
	return dbs.Orm.Raw(s.SQL, s.Args...)
}

func (s *SQLMapper) Exec() *gorm.DB {
	if s.db != nil {
		return s.db.Exec(s.SQL, s.Args...)
	}
	return dbs.Orm.Exec(s.SQL, s.Args...)
}

func Mapper(sql string, args []interface{}, err error) *SQLMapper {
	if err != nil {
		panic(err.Error())
	}
	return NewSQLMapper(sql, args)
}
