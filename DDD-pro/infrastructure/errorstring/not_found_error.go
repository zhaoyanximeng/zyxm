package errorstring

import "fmt"

const (
	NotFoundRootID = 1404 		// 没有设置根实体ID
	NotFoundModelData = 2404	// 实体数据没有找到
)

type NotFoundError struct {
	Code int
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(code int,message string) *NotFoundError {
	return &NotFoundError{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundIDError(modelName string) *NotFoundError {
	return &NotFoundError{
		Code:    NotFoundRootID,
		Message: fmt.Sprintf("model %s ID not found",modelName),
	}
}

func NewNotFoundDataError(modelName,err string) *NotFoundError {
	return &NotFoundError{
		Code:    NotFoundModelData,
		Message: fmt.Sprintf("model %s ID not found %s",modelName,err),
	}
}