package Errors

import "fmt"

const (
	DataError_ModelCreate = 1501 // 实体新增失败
	DataError_ModelUpdate = 1502 // 实体更新失败
)

type DataOperatorError struct {
	Code    int
	Message string
}

func (s *DataOperatorError) Error() string {
	return s.Message
}
func NewDataOperatorError(code int, message string) *DataOperatorError {
	return &DataOperatorError{Code: code, Message: message}
}
func NewModelCreateError(modelName string, err string) *DataOperatorError {
	return &DataOperatorError{Code: DataError_ModelCreate,
		Message: fmt.Sprintf("Model %s create error:%s", modelName, err),
	}
}
