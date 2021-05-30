package result

import "fmt"

type ErrorResult struct {
	data interface{}
	err  error
}

func (s *ErrorResult) Unwrap() interface{} {
	if s.err != nil {
		panic(s.err.Error())
	}
	return s.data
}

func Result(vs ...interface{}) *ErrorResult {
	if len(vs) == 1 {
		if vs[0] == nil {
			return &ErrorResult{data: nil, err: nil}
		}
		if e, ok := vs[0].(error); ok {
			return &ErrorResult{data: nil, err: e}
		}
	}
	if len(vs) == 2 {
		if vs[1] == nil {
			return &ErrorResult{data: vs[0], err: nil}
		}
		if e, ok := vs[1].(error); ok {
			return &ErrorResult{data: vs[0], err: e}
		}
	}
	return &ErrorResult{data: nil, err: fmt.Errorf("result format error")}
}
