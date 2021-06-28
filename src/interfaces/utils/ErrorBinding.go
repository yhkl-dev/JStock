package utils

type ErrorResult struct {
	data interface{}
	err error
}
func NewErrorResult(data interface{}, err error) *ErrorResult {
	return &ErrorResult{data: data, err: err}
}
func(this *ErrorResult) Unwrap() interface{}{
	if this.err!=nil{
		panic(this.err.Error())
	}
	return this.data
}

type BindFunc func(v interface{}) error
func Exec(f BindFunc,value interface{}) *ErrorResult  {
		err:=f(value)
		return NewErrorResult(value,err)

}