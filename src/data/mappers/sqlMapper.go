package mappers

type SQLMapper struct {
	SQL  string
	Args []interface{}
}

func NewSQLMapper(sql string, args []interface{}) *SQLMapper {
	return &SQLMapper{SQL: sql, Args: args}
}

func Mapper(sql string, args []interface{}, err error) *SQLMapper {
	if err != nil {
		panic(err.Error())
	}
	return NewSQLMapper(sql, args)
}
