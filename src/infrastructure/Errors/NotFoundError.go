package Errors

import "fmt"

const (
	NotFound_RootID    = 1401 //没有设置 根实体的ID
	NotFound_ModelData = 1402 // 实体数据没有找到
)

type NotFoundError struct {
	Code    int
	Message string
}

func (s *NotFoundError) Error() string {
	return s.Message
}
func NewNotFoundError(code int, message string) *NotFoundError {
	return &NotFoundError{Code: code, Message: message}
}
func NewNotFoundIDError(modelName string) *NotFoundError {
	return &NotFoundError{
		Code:    NotFound_RootID,
		Message: fmt.Sprintf("Model %s ID not found", modelName),
	}
}
func NewNotFoundDataError(modelName string, err string) *NotFoundError {
	return &NotFoundError{
		Code:    NotFound_ModelData,
		Message: fmt.Sprintf("Model %s data not found %s", modelName, err),
	}
}
