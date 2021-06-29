package utils

type ErrorResult struct {
	data interface{}
	err  error
}

func NewErrorResult(data interface{}, err error) *ErrorResult {
	return &ErrorResult{data: data, err: err}
}
func (s *ErrorResult) Unwrap() interface{} {
	if s.err != nil {
		panic(s.err.Error())
	}
	return s.data
}

type BindFunc func(v interface{}) error

func Exec(f BindFunc, value interface{}) *ErrorResult {
	err := f(value)
	return NewErrorResult(value, err)
}
